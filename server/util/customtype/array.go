package customtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

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
