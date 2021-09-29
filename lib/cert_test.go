package lib

import "testing"

func TestGetCertList(t *testing.T) {
	ret, err := GetCert("prod.app.hengyangtai.xyz")
	if err != nil {
		panic(err)
	}
	for _, item := range ret {
		t.Log(item)
	}
}
