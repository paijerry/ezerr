package ezerr

import (
	"fmt"
	"testing"
)

func TestEzError(t *testing.T) {
	err := GenEzErr("100", "message")

	ok := IsEz(err)

	if !ok {
		t.Fatal("error is not MyError")
	}

	if err.Error() != "100: message" {
		t.Fatal("message error")
	}

	fmt.Println(err)
	fmt.Println(err.GetCode())
	fmt.Println(err.GetMsg())

	err.SetDetail("hello")
	err.SetDetail("hello2")
	fmt.Println(err.GetDetail())
}
