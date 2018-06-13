package ezerr

import (
	"fmt"
	"testing"
)

func TestEzError(t *testing.T) {
	err := New("message", "100")

	ok := IsEz(err)

	if !ok {
		t.Fatal("error is not MyError")
	}

	if err.Error() != "<100> message" {
		t.Fatal("message error")
	}

	fmt.Println(err)
	fmt.Println(err.GetCode())
	fmt.Println(err.GetMsg())

	err.SetDetail("hello1").SetDetail(2)
	err.SetDetail("hello3")
	fmt.Println(err.GetDetail())
	fmt.Println(err.Error())
}
