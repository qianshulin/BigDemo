package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	app.Get("/", demo)
	app.Listen(":8080")
}

func demo(context.Context) {
	fmt.Println("你好世界")
}
