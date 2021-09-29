package lib

import (
	"fmt"
	"math"
	"time"
)

const DayLayout = "2006-01-02"
const DateTimeLayout = "2006-01-02 15:04:05"

func TodayStartTime() time.Time {
	nowTime := time.Now()
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
}

func SecondsDiffFromToday(t time.Time) int64 {
	return SecondsDiff(TodayStartTime(), t)
}

func SecondsDiff(from time.Time, to time.Time) int64 {
	return to.Unix() - from.Unix()
}

func ParseDateTime(s string) (time.Time, error) {
	return time.ParseInLocation(DateTimeLayout, s, time.Local)
}

func ParseDayTime(s string) (time.Time, error) {
	return time.ParseInLocation(DayLayout, s, time.Local)
}

func ParseTime(s string) (time.Time, error) {
	return time.ParseInLocation("15:06:05", s, time.Local)
}

func NowTimePointer() *time.Time {
	var t time.Time
	t = time.Now()
	return &t
}

func SimplifyTimeFormat(t time.Time) string {
	n := time.Now()
	d := n.Sub(t)
	suffix := "前"
	if d < 0 {
		suffix = "后"
	}
	if d < time.Minute {
		return fmt.Sprintf("%d秒"+suffix, int(math.Abs(d.Seconds())))
	} else if d < time.Hour {
		return fmt.Sprintf("%d分钟"+suffix, int(math.Abs(d.Minutes())))
	} else if t.Format(DayLayout) == n.Format(DayLayout) {
		return t.Format("15:04")
	}
	return t.Format(DateTimeLayout)
}

/**
 * 计算函数执行时间
 */
func RunDuration(runner func()) time.Duration {
	start := time.Now()
	runner()
	return time.Now().Sub(start)
}
