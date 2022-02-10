package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// reference:
// https://mojotv.cn/2019/07/30/golang-http-request

func main() {
	// r, err := http.Get("https://api.github.com/events")
	r, err := http.Get("https://qq.com")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", body)
}
