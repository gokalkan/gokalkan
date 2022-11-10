package ckalkan

import "fmt"

// Close закрывает связь с динамической библиотекой.
func (cli *Client) Close() (err error) {
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

	cli.XMLFinalize()
	cli.Finalize()

	err = cli.handler.close()
	if err != nil {
		return err
	}

	return nil
}
