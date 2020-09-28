package util

import "github.com/Treblex/go-echo-demo/server/util/mlog"

// Recover 使用defer调用阻止painc中止程序
func Recover() {
	defer func() {
		if r := recover(); r != nil {
			mlog.Printf("recover拦截:%v\n\n", r)
		}
	}()
}
