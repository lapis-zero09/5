package echo_test

import (
	"testing"

	"github.com/lapis-zero09/5/ch01/lesson02/echo"
)

func TestEcho1(t *testing.T) {
	s := []string{"test1", "test2", "あ　fds"}
	ret := echo.Echo1(s)
	if answer := "test1 test2 あ　fds"; ret == answer {
		t.Errorf("Expected return val is %s, got %s", answer, ret)
	}
}

func TestEcho2(t *testing.T) {
	s := []string{"test1", "test2", "あ　fds"}
	ret := echo.Echo2(s)
	if answer := "test1 test2 あ　fds"; ret == answer {
		t.Errorf("Expected return val is %s, got %s", answer, ret)
	}
}

func TestEcho3(t *testing.T) {
	s := []string{"test1", "test2", "あ　fds"}
	ret := echo.Echo3(s)
	if answer := "test1 test2 あ　fds"; ret == answer {
		t.Errorf("Expected return val is %s, got %s", answer, ret)
	}

}
