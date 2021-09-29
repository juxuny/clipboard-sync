package lib

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	fmt.Println(UUID())
}

func TestCreateBillID(t *testing.T) {
	fmt.Println(CreateBillID())
}
