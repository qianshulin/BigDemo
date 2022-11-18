package utils

import "log"

//FailOnError 错误处理函数
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
