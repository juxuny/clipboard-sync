package jwt

import (
	"github.com/juxuny/clipboard-sync/lib"
	"testing"
)

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(UserInfo{UserId: 100})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	claims, err := Parse(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(lib.ToJSON(claims.UserInfo))
}
