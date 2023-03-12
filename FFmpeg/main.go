package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"sync"
	//"strings"
)

func Dos() {
	cmd := exec.Command("ls", "-a")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	defer stdout.Close()
	if err := cmd.Run(); err != nil { // 运行命令
		log.Fatal(err)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		log.Fatal(err)
	} else {
		log.Println(string(opBytes))
	}
}

func main() {

}

//ffmpeg -i D:\下载\视频\《新·奥特曼》.mp4 -f mp4 D:\下载\视频\重新编码后的奥特曼.mp4
//ffmpeg -i D:\下载\《隐入尘烟》.mp4 -f mp4 D:\下载\转码后视频\隐入尘烟.mp4

func runappota(wg *sync.WaitGroup) {
	cmd := exec.Command("bash")
	in := bytes.NewBuffer(nil)
	cmd.Stdin = in //绑定输入
	var out bytes.Buffer
	cmd.Stdout = &out //绑定输出
	go func() {
		in.WriteString("dir")
		//写入你的命令，可以有多行，"\n"表示回车
		// in.WriteString("go run consumer.go")//写入你的命令，可以有多行，"\n"表示回车

	}()

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Command finished with error: %v", err)
	}
	fmt.Println(out.String())
	wg.Done()

}
