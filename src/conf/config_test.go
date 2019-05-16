package conf

import (
	"fmt"
	"testing"
)

func TestAgfunInst(t *testing.T) {
	fmt.Println(AgfunInst().VideoHost)
	if len(AgfunInst().VideoHost) > 0 {
		t.Log("success")
	} else {
		t.Error("fail")
	}
}
