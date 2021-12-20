package kalkan

import "errors"

// returnErr возвращает последнюю глобальную ошибку, если returnCode не равен 0
func (cli *Client) returnErr(returnCode int) error {
	if returnCode != 0 {
		return errors.New(cli.GetLastErrorString())
	}
	return nil
}
