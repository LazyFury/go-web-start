package customtype

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// NumberTime 时间戳
type NumberTime struct {
	time.Time
}

// UnmarshalJSON UnmarshalJSON
func (t *NumberTime) UnmarshalJSON(b []byte) error {
	var str = string(b)
	num, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	t.Time = time.Unix(int64(num/1000), 0)
	return nil
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
