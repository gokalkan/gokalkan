package ckalkan

import (
	"testing"
)

func TestGetErrorCode(t *testing.T) {
	er := KalkanError{
		errorCode:   ErrorCodeCertParseError,
		errorString: "test",
	}

	got, ok := GetErrorCode(er)
	if !ok {
		t.Fatal("expected true")
	}

	if got != ErrorCodeCertParseError {
		t.Fatal("mismatch eror code")
	}

	er2 := &KalkanError{
		errorCode:   ErrorCodeCheckChainError,
		errorString: "test",
	}

	got, ok = GetErrorCode(er2)
	if !ok {
		t.Fatal("expected true")
	}

	if got != ErrorCodeCheckChainError {
		t.Fatal("mismatch eror code")
	}
}
