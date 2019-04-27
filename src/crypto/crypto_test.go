package crypto

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	str1 := "qwertyqqqwee"
	str2 := "uiopssssssffsfsff"
	md5 := MD5(str1)
	s := MD5(str2)
	if len(md5) == len(s) {
		fmt.Println(len(md5))
		t.Log("pass")
	} else {
		t.Error("unpass")
	}
}

func TestSHA1(t *testing.T) {
	str1 := "qwertyqqqwee"
	sha1 := SHA1(str1)
	fmt.Println(sha1)
}
