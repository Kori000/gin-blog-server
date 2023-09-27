package main

import (
	"fmt"
	"gin-blog-server/core"
	"gin-blog-server/global"
)

func main() {
	core.InitCore()
	fmt.Print(global.Config)
}
