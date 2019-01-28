package etcddb

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	client := GetCli()
	a := A{
		Name: "xiaohua",
		Age:  15,
	}
	e := client.Put("mykey", a)
	if e != nil {
		t.Error("unpass")
	}

	var b A
	e = client.Get("mykey", &b)
	if e != nil {
		t.Error("unpass")
	}

	fmt.Println("1111111111111111", b)
}

type A struct {
	Name string
	Age  int
}
