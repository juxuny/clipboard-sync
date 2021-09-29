package lib

import (
	"testing"
	"time"
)

func TestSimplifyTimeFormat(t *testing.T) {
	var data = []struct {
		time   time.Time
		result string
	}{
		{time: time.Now().Add(-time.Second * 4), result: "4秒前"},
		{time: time.Now().Add(-time.Minute * 4), result: "4分钟前"},
		{time: time.Now().Add(-time.Hour), result: time.Now().Add(-time.Hour).Format("15:04")},
		{time: time.Now().Add(-24 * time.Hour), result: time.Now().Add(-24 * time.Hour).Format(DateTimeLayout)},
	}
	for _, item := range data {
		s := SimplifyTimeFormat(item.time)
		if s != item.result {
			t.Fatal("wrong result: ", s)
		}
		t.Log(s)
	}
}
