package util

import (
	"fmt"
)

// Recover 使用defer调用阻止painc中止程序
func Recover() {
	if r := recover(); r != nil {
		fmt.Printf("recover拦截:%v\n\n", r)
	}
}
