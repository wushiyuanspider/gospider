package main

import (
	"bufio"
	"fmt"
	"gospider/src/controller"
	"os"
)

func main() {
	var filepath string
	reader := bufio.NewReader(os.Stdin)

	if len(os.Args) != 2 {
		fmt.Printf("Please input ini file path: ")
		confPath, _, _ := reader.ReadLine()
		filepath = string(confPath)
	} else {
		filepath = os.Args[1]
	}

	// 初始化控制器,并打印信息
	controller.Init(filepath, true)

	fmt.Printf("==========\nStart? (y/n):  ")
	buffer, _, _ := reader.ReadLine()
	if ok := string(buffer); ok == "y" || ok == "Y" || ok == "yes" {
		fmt.Printf("开始爬取配置文件 %s 指定内容......\n", filepath)
		controller.Run()
	} else {
		return
	}
}
