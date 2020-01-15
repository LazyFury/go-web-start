package util

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// TimeFormat 自定义时间格式化
func TimeFormat(t time.Time, arg ...string) string {
	str := "y-m-d h:i:s"
	if len(arg) > 0 {
		str = arg[0]
	}
	result := ""
	for _, ch := range str {
		s := string(ch)
		switch s {
		case "y":
			result += "2006"
			break
		case "m":
			result += "01"
			break
		case "d":
			result += "02"
			break
		case "h":
			result += "15"
			break
		case "i":
			result += "04"
			break
		case "s":
			result += "05"
			break
		default:
			result += s
			break
		}
	}

	return t.Format(result)
}

// ItoaFix0 ItoaFix0
func ItoaFix0(i int) string {
	if i < 10 {
		return fmt.Sprintf("0%d", i)
	}
	return fmt.Sprintf("%d", i)
}

// LocalTime LocalTime
type LocalTime struct {
	time.Time
}

var timeLayout string = "2006年01月02日 15:04:05"

// MarshalJSON json格式化时间的方法
func (t LocalTime) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(timeLayout)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, timeLayout)
	b = append(b, '"')
	return b, nil
}

// Value Value
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan Scan
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
