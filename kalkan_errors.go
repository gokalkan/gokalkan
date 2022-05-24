package gokalkan

import (
	"errors"
	"fmt"
)

var (
	ErrPanic = errors.New("panic recovered")
)

type KalkanError struct {
	errorCode   KCErrorCode
	errorString string
}

func (e KalkanError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.errorCode.Hex(), e.errorCode.String(), e.errorString)
}

// GetErrorCode извлекает из ошибки KCErrorCode из ошибка типа KalkanError.
func GetErrorCode(err error) (code KCErrorCode, ok bool) {
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
