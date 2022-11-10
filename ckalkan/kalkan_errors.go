package ckalkan

import (
	"errors"
	"fmt"
)

var ErrPanic = errors.New("panic recovered")

type KalkanError struct {
	errorCode   ErrorCode
	errorString string
}

func (e KalkanError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.errorCode.Hex(), e.errorCode.String(), e.errorString)
}

// GetErrorCode извлекает из ошибки ErrorCode из ошибка типа KalkanError.
func GetErrorCode(err error) (code ErrorCode, ok bool) {
	var kalkanErr KalkanError
	if errors.As(err, &kalkanErr) {
		return kalkanErr.errorCode, true
	}

	var kalkanErr2 *KalkanError
	if errors.As(err, &kalkanErr2) {
		return kalkanErr2.errorCode, true
	}

	return 0, false
}

// wrapError возвращает последнюю глобальную ошибку, если returnCode не равен 0
func (cli *Client) wrapError(returnCode int) error {
	if returnCode != 0 {
		ec, es := cli.GetLastErrorString()

		return KalkanError{
			errorCode:   ec,
			errorString: es,
		}
	}

	return nil
}
