package main

import (
	"flag"
	"fmt"
	"os"
)

// Go语言-命令行参数（os.Args, flag包）
// https://blog.csdn.net/guanchunsheng/article/details/79612153

func main() {
	cmdline := os.Args
	fmt.Println(cmdline)

	args := os.Args[1:]
	fmt.Println(args)

	verbose := flag.Bool("v", false, "verbose output")
	name := flag.String("name", "jonny", "name input")

	flag.Parse()

	fmt.Println(*verbose, *name)

	notParse := flag.Args()
	fmt.Println(notParse)
}
