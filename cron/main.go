package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	//crontab := cron.New()  默认从分开始进行时间调度
	crontab := cron.New(cron.WithSeconds()) //精确到秒
	//定义定时器调用的任务函数
	a := func() {
		fmt.Println("hello world", time.Now())
	}
	b := func() {
		fmt.Println("你好世界", time.Now())
	}

	//定时任务
	//spec := "0 0 0 1/1 * ?" //cron表达式，每天0点0分执行
	spec1 := "0/3 * * * * ?"
	spec2 := "0/6 * * * * ?"
	// 添加定时任务,
	crontab.AddFunc(spec1, a)
	crontab.AddFunc(spec2, b)

	// 启动定时器
	crontab.Start()
	// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	select {} //阻塞主线程停止
}
