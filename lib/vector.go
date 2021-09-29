package lib

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Vector []float64

func NewVector(num int, x float64) Vector {
	ret := make(Vector, num)
	for i := range ret {
		ret[i] = x
	}
	return ret
}

func (t Vector) Add(x Vector) (ret Vector, err error) {
	if len(t) != len(x) {
		return ret, fmt.Errorf("the length of two vector is not equal %d != %d", len(t), len(x))
	}
	ret = make(Vector, len(t))
	for i := 0; i < len(ret); i++ {
		ret[i] = t[i] + x[i]
	}
	return ret, nil
}

func (t Vector) DrawOne() (index int, err error) {
	if len(t) == 0 {
		return -1, errors.Errorf("invalid vector, len = %d", len(t))
	}
	for i := 0; i < len(t); i++ {
		if t[i] >= 1 {
			return i, nil
		}
	}
	r := math.Abs(rand.Float64())
	s := float64(0)
	index = 0
	for i := 0; i < len(t); i++ {
		s += t[i]
		if i == len(t)-1 {
			if r <= s {
				return i, nil
			} else {
				return -1, ErrNotFound
			}
		}
		if r > s {
			continue
		}
		return i, nil
	}
	return -1, ErrNotFound
}

type BigVector []*big.Float

func (t Vector) BigVector() BigVector {
	ret := make(BigVector, len(t))
	for i, item := range t {
		ret[i] = big.NewFloat(item)
	}
	return ret
}

func (t Vector) Reset(num int) Vector {
	if num <= 0 {
		return t
	}
	for i := 0; i < num && i < len(t); i++ {
		t[i] = 0
	}
	return t
}
