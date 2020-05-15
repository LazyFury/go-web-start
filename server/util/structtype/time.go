package structtype

import (
	"database/sql/driver"
	"fmt"
	"time"
)

var (
	// DefaultTimeLayout DefaultTimeLayout
	DefaultTimeLayout string = "2006年01月02日 15:04:05"
)

type (
	// LocalTime LocalTime
	LocalTime struct {
		time.Time
	}
	// CustomTime CustomTime
	CustomTime LocalTime
	// TestTime TestTime
	TestTime struct {
		Atime CustomTime
		Btime LocalTime
	}
)

func (t *CustomTime) localLayout() string {
	return "2006年01月02------- 15:04:05"
}

func (t *LocalTime) localLayout() string {
	return "2006年01月02日 15:04:05"
}

// MarshalJSON json格式化时间的方法
func (t LocalTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t.Time).Format(t.localLayout()))
	return []byte(stamp), nil
}

// UnmarshalJSON UnmarshalJSON
func (t *LocalTime) UnmarshalJSON(b []byte) error {
	var str = string(b)
	fmt.Printf(str + "\n")
	tTime, err := time.Parse(`"`+t.localLayout()+`"`, str)
	if err != nil {
		fmt.Println(err)
	}
	t.Time = tTime
	return nil
}

// Value Value
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
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

// NumberTime 时间戳
type NumberTime struct {
	time.Time
}

// MarshalJSON json格式化时间的方法
func (t NumberTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("%d", time.Time(t.Time).Unix())
	return []byte(stamp), nil
}

// Value Value
func (t NumberTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan Scan
func (t *NumberTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = NumberTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
