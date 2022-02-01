package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func checkErr(e error) {
	if e != nil {
		log.Fatalln("error!!!")
	}
}

func main() {

	pwd, e := os.Getwd()
	checkErr(e)
	fmt.Printf("get path: %s\n", pwd)

	fpath := pwd + "/godemos/data/io.txt"

	content, err := ioutil.ReadFile(fpath)
	checkErr(err)

	fmt.Println(string(content))
}
