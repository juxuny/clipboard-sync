package lib

import "regexp"

type checker struct{}

func Checker() *checker {
	return &checker{}
}

func (*checker) IsPhone(f string) (b bool) {
	b, e := regexp.MatchString("^(1)\\d{10}$", f)
	if e != nil {
		b = false
	}
	return
}

func (*checker) IsPassword(s string) (b bool) {
	b, e := regexp.MatchString("^[0-9A-Za-z@\\-_.]{6,32}$", s)
	if e != nil {
		b = false
	}
	return
}

func (*checker) IsUserName(s string) (b bool) {
	b, e := regexp.MatchString("^[0-9A-Za-z@\\-_.]{4,16}$", s)
	if e != nil {
		b = false
	}
	return
}
