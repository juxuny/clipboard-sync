package lib

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	fmt.Println(MD5("000000"))
}

func TestCheckPasswordLever(t *testing.T) {
	fmt.Println(CheckPasswordLever("@Aa123456"))
}
