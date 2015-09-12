package main

import (
    "fmt"
    //"time"
    //"path"
    "os"
    "bufio"
)

var spider *Spider

func main() {
    var err error
    reader := bufio.NewReader(os.Stdin)
    
    if len(os.Args) != 2 {
        fmt.Printf("Please input ini file path: ")
        confPath, _, _ := reader.ReadLine()
        spider, err = NewSpider(string(confPath))
    } else {
        spider, err = NewSpider(os.Args[1])
    }

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Name:\t", spider.name)
        fmt.Println("URL:\t", spider.startURL)
        fmt.Println("depth:\t", spider.depth)
    }
    
    fmt.Printf("Start? (y/n):  ")
    buffer, _, _ := reader.ReadLine()
    if  ok := string(buffer); ok == "y" || ok == "Y" || ok == "yes" {
        fmt.Printf("开始爬取 %s 的指定内容......\n", spider.startURL)
    } else {
        return
    }
}


