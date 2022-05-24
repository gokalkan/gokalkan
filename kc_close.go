package gokalkan

import "fmt"

// KCClose закрывает связь с динамической библиотекой
func (cli *KCClient) KCClose() (err error) {
	cli.mu.Lock()
	defer cli.mu.Unlock()

	defer func() {
		if r := recover(); r != nil {
			if err != nil {
				err = fmt.Errorf("%w: %s", err, r)
			} else {
				err = fmt.Errorf("%w: %s", ErrPanic, r)
			}
		}
	}()

	cli.KCXMLFinalize()
	cli.KCFinalize()

	err = cli.handler.Close()
	if err != nil {
		return err
	}

	return nil
}
