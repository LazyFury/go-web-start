package customtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// JSONObject json => sql text field
type JSONObject struct {
	JSON map[string]interface{} `gorm:"type:text"`
}

// UnmarshalJSON UnmarshalJSON
func (j *JSONObject) UnmarshalJSON(b []byte) error {
	j.JSON = map[string]interface{}{}
	if err := json.Unmarshal(b, &j.JSON); err != nil {
		return err
	}
	return nil
}

// MarshalJSON MarshalJSON
func (j *JSONObject) MarshalJSON() (b []byte, err error) {
	return json.Marshal(&j.JSON)
}

// Value Value
func (j JSONObject) Value() (driver.Value, error) {
	if j.JSON == nil {
		return nil, nil
	}
	b, err := json.Marshal(&j.JSON)
	if err != nil {
		return nil, err
	}
	str := string(b)
	if str == "" {
		return nil, err
	}
	return str, nil
}

// Scan Scan
func (j *JSONObject) Scan(v interface{}) error {
	b, ok := v.([]byte)
	if ok {
		if len(b) == 0 {
			b = []byte("{}")
		}
		return j.UnmarshalJSON(b)
	}
	return fmt.Errorf("类型错误:%v", reflect.TypeOf(v))
}
