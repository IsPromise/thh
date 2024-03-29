package test

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(_ *testing.T) {
	zeroTime := time.Time{}
	zeroTime2 := time.Time{}
	fmt.Println(zeroTime.Format(time.RFC3339))
	fmt.Println(zeroTime == zeroTime2)
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixMicro())
	fmt.Println(now.UnixNano())
	fmt.Println(now.Nanosecond())
	year, month, day := now.Date()
	fmt.Println(year, month, day)
	hour, minute, second := now.Clock()
	fmt.Println(hour, minute, second)
	fmt.Println(now.YearDay())
	// Y-m-d H:i:s
	// 2006-01-02 03:04:05
	// 2006-01-02 15:04:05
	// 2006,1,2,3(15),4,5
	fmt.Println(now.Format("20060102030405"))
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	layout := "2006-01-02 15:04:05"
	fmt.Println(time.Unix(0, 0).Format(layout))
	t, _ := time.ParseInLocation(layout, "2011-01-01 20:20:20", time.Local)
	fmt.Println(t.Format(layout))
	times := "2020-09-18 15:04:05"
	tagStartTime, _ := time.Parse("2025-01-01 01:01:01", times)
	tagEndTime, _ := time.Parse("2025-12-31 01:01:01", times)
	fmt.Println(tagStartTime, tagEndTime)
}
