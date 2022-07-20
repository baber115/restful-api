package main

import (
	"fmt"
	"go-restful-api/cmd"
)

// 启动接口程序
// go run main.go start
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
