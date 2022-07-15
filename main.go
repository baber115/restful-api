package main

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/go-restful-api/cmd"
	"fmt"
)

// 启动接口程序
// go run main.go start
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
