package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// reference:
// https://mojotv.cn/2019/07/30/golang-http-request

var httpURL string = "https://baidu.com"

func httpGet() {
	resp, err := http.Get(httpURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%s", body)
	fmt.Printf("%#v\n", resp)
}

func httpGetWithParam() {
	req, err := http.NewRequest(http.MethodGet, httpURL, nil)
	if err != nil {
		panic(err)
	}

	params := make(url.Values)
	params.Add("key1", "value1")
	params.Add("key2", "value2")

	req.URL.RawQuery = params.Encode()
	fmt.Println(req.URL.String())
	// output
	// https://baidu.com?key1=value1&key2=value2

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf("%s", body)
	fmt.Printf("%#v\n", resp)
}

func main() {
	httpGet()
	// httpGetWithParam()
}
