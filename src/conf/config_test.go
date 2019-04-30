package conf

import (
	"fmt"
	"testing"
)

func TestAgfunInst(t *testing.T) {
	fmt.Println(AgfunInst().Etcd)
	if len(AgfunInst().Etcd) > 0 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
