package test

import (
	"fmt"
	"os"
	"testing"
)

//Print  输出
func Test_Print(t *testing.T) {
	fmt.Print("直接输出内容", "\n") // \n :手动换行
	fmt.Println("自动加换行")
	name := "罗时瑞"
	fmt.Printf("我是：%s\n", name)
}

func Test_Fprint(t *testing.T) {
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xxx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	//name := "罗时瑞"
	// 向打开的文件句柄中写入内容
	//fmt.Fprintf(fileObj, "往文件中写入内容：%s", name)
	fmt.Fprintln(fileObj, "向标准输出写入内容")
}

func Test_Sprint(t *testing.T) {
	s1 := fmt.Sprint("罗时瑞")
	name := "罗时瑞"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("罗时瑞")
	s4 := fmt.Sprintln("罗时瑞")
	fmt.Println(s1, s2, s3, s4)
}

func Test_Errorf(t *testing.T) {
	err := fmt.Errorf("错误码:%d", 404)
	fmt.Println("发生一个错误，", err)
}
