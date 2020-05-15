package structtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

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
