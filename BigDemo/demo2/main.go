package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	//得到请求信息req
	req, _ := http.NewRequest("GET", "http://httpbin.org/get", nil)
	//设置请求头
	req.Header.Add("name", "lyizriii")
	req.Header.Add("age", "22")
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("string(b): %v\n", string(b))

}
