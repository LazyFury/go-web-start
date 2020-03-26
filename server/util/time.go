package util

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

var (
	// DefaultTimeLayout DefaultTimeLayout
	DefaultTimeLayout string = "2006年01月02日 15:04:05"
)

// LocalTime LocalTime
type LocalTime struct {
	time.Time
}

func (t LocalTime) localLayout() string {
	return "2006-01月02日 15:04:05"
}

// MarshalJSON json格式化时间的方法
func (t LocalTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t.Time).Format(t.localLayout()))
	return []byte(stamp), nil
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

// NumberTime NumberTime
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

// Money Moneytype StrictFloat64 float64
type Money struct {
	Money float64
}

// MarshalJSON json格式化时间的方法
func (money Money) MarshalJSON() ([]byte, error) {
	// fmt.Printf("step jsonMarsha\n")
	var stamp = fmt.Sprintf("\"%.2f\"", money.Money)
	return []byte(stamp), nil
}

// Scan Scanner
func (money *Money) Scan(v interface{}) error {
	value, ok := v.([]uint8)
	if ok {
		val, err := strconv.ParseFloat(string(value), 64)
		if err != nil {
			return err
		}
		*money = Money{Money: val}
		return nil
	}
	return fmt.Errorf("%v 类型错误  scan失败", reflect.TypeOf(v))
}
