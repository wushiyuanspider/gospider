package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)


func StartSearch(spider *Spider) {
	res, err := http.Get(spider.startURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}