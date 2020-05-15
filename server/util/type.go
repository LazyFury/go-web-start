package util

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type gormType interface {
	Value() (driver.Value, error)
	Scan(v interface{}) error
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

var (
	// DefaultTimeLayout DefaultTimeLayout
	DefaultTimeLayout string = "2006年01月02日 15:04:05"
)

// LocalTime LocalTime
type LocalTime struct {
	time.Time
}

func (t LocalTime) localLayout() string {
	return "2006年01月02日 15:04:05"
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

// Money Moneytype StrictFloat64 float64
type Money float64

// MarshalJSON json格式化时间的方法
func (money Money) MarshalJSON() ([]byte, error) {
	// fmt.Printf("step jsonMarsha\n")
	var stamp = fmt.Sprintf("\"%.2f\"", money)
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
		*money = Money(val)
		return nil
	}
	return fmt.Errorf("%v 类型错误  scan失败", reflect.TypeOf(v))
}

// Value Value
func (money *Money) Value() (driver.Value, error) {
	return float64(*money), nil
}

// UnmarshalJSON UnmarshalJSON
func (money *Money) UnmarshalJSON(b []byte) error {
	// b = bytes.Trim(b, "\"")
	// fmt.Printf("%v", string(b))

	var num float64
	if err := json.Unmarshal(b, &num); err != nil {
		return err
	}
	*money = Money(num)
	// fmt.Printf("Array UnmarshalJSON %v \n", arr)
	return nil
}

// Array json传数组类型
type Array []string

// UnmarshalJSON req.body []byte=>对象，记得调用json.Unmarshal要新建原始类型进行绑定，不如会死循环
func (a *Array) UnmarshalJSON(b []byte) error {
	// b = bytes.Trim(b, "\"")
	// fmt.Printf("%v", string(b))

	arr := []string{}
	if err := json.Unmarshal(b, &arr); err != nil {
		return err
	}
	*a = Array(arr)
	// fmt.Printf("Array UnmarshalJSON %v \n", arr)
	return nil
}

// Value 存库,对象到转储数据 标准字符 int类型
func (a Array) Value() (driver.Value, error) {
	// fmt.Printf("value %v \n", a)
	if len(a) == 0 {
		return nil, nil
	}
	arr := strings.ReplaceAll(strings.Trim(fmt.Sprint(a), "[]"), " ", ",")
	return arr, nil
}

// Scan 绑定，数据库到对象,这里到数据取到到都是[]uint8字节，转化为对象
func (a *Array) Scan(v interface{}) error {
	// fmt.Printf("scan")
	value, ok := v.([]uint8)
	if ok {
		arr := strings.Split(string(value), ",")
		*a = Array(arr)
		return nil
	}
	return fmt.Errorf("%v 类型错误  scan失败", reflect.TypeOf(v))
}

// MarshalJSON 对象到json转换 接口展示
func (a *Array) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal([]string(*a))
	// fmt.Printf("MarshalJSON %v \n", string(b))
	if err != nil {
		return nil, err
	}
	// 空数组默认值
	if string(b) == "null" {
		b = []byte(`[]`)
	}
	return b, nil
}
