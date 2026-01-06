package gokalkan

import (
	"log/slog"
	"sync"
)

// ClientPool представляет пул переиспользуемых клиентов Kalkan.
// Вместо создания и закрытия клиентов при каждом использовании,
// ClientPool переиспользует существующие клиенты, что минимизирует
// количество Init/Finalize циклов и снижает утечки памяти в C коде.
type ClientPool struct {
	pool    sync.Pool
	opts    []Option
	mu      sync.Mutex
	created int
	stored  int
}

// NewClientPool создает новый пул клиентов с заданными опциями.
//
// Пример использования:
//
//	pool := gokalkan.NewClientPool(gokalkan.OptsProd...)
//	cli, err := pool.Get()
//	if err != nil {
//	    return err
//	}
//	defer pool.Put(cli) // НЕ cli.Close()!
//
//	// Используем клиент
//	signedXML, err := cli.SignXML("<root>data</root>")
func NewClientPool(opts ...Option) *ClientPool {
	cp := &ClientPool{opts: opts}
	cp.pool.New = func() interface{} {

		cp.mu.Lock()
		defer cp.mu.Unlock()
		cp.created += 1

		cli, err := NewClient(opts...)
		if err != nil {
			// Возвращаем nil при ошибке, Get() обработает это
			return nil
		}
		return cli
	}
	return cp
}

// Get получает клиент из пула или создает новый если пул пуст.
// Клиент должен быть возвращен в пул через Put() после использования.
func (p *ClientPool) Get() (*Client, error) {
	obj := p.pool.Get()
	if obj == nil {
		slog.Info(
			"pool: не смогли получить, создаем экземпляр в пул",
			slog.Int("stored", p.stored),
			slog.Int("created", p.created),
		)
		// pool.New вернула nil из-за ошибки, создаем напрямую
		return NewClient(p.opts...)
	}

	cli, ok := obj.(*Client)
	if !ok || cli == nil {
		slog.Info(
			"pool: некорректный объект, создаем",
			slog.Int("stored", p.stored),
			slog.Int("created", p.created),
		)
		// Некорректный объект в пуле, создаем новый
		return NewClient(p.opts...)
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	p.stored -= 1

	slog.Info(
		"pool: получили объект из пула",
		slog.Int("stored", p.stored),
		slog.Int("created", p.created),
	)

	return cli, nil
}

// Put возвращает клиент обратно в пул для переиспользования.
// ВАЖНО: НЕ вызывайте cli.Close() на клиентах из пула!
// Вместо этого всегда используйте Put().
func (p *ClientPool) Put(cli *Client) {
	if cli == nil {
		return
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.stored += 1
	p.created -= 1

	slog.Info(
		"pool: получили пустой экземпляр",
		slog.Int("stored", p.stored),
		slog.Int("created", p.created),
	)

	// Возвращаем клиент в пул БЕЗ вызова Close()
	p.pool.Put(cli)
}

// Close закрывает пул и все активные клиенты в нем.
// После вызова Close() пул больше не может использоваться.
//
// ВНИМАНИЕ: Эта функция НЕ может закрыть клиенты которые в данный момент
// находятся вне пула (полученные через Get() но не возвращенные через Put()).
// Пользователь должен сам позаботиться о возврате всех клиентов перед Close().
func (p *ClientPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	slog.Info(
		"pool: количество объектов",
		slog.Int("stored", p.stored),
		slog.Int("created", p.created),
	)

	// Создаем временный слайс для клиентов
	// Извлекаем и закрываем все клиенты из пула
	var clients []*Client
	for {
		obj := p.pool.Get()
		if obj == nil {
			break
		}
		if cli, ok := obj.(*Client); ok && cli != nil {
			clients = append(clients, cli)
		}
	}

	// Закрываем всех извлеченных клиентов
	var lastErr error
	for _, cli := range clients {
		if err := cli.Close(); err != nil {
			lastErr = err
		}
	}

	return lastErr
}
