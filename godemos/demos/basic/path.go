package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func execPath() {
	cpath, _ := os.Getwd()
	log.Println("current path:", cpath)

	epath, _ := exec.LookPath(os.Args[0])
	log.Println("exec bin file path:", epath)

	dir, exe := path.Split(epath)
	log.Println("exec bin file dir:", dir)
	log.Println("exec bin file name:", exe)

	cpath, _ = os.Getwd()
	log.Println("current path:", cpath)
	os.Chdir(dir)
	wd, _ := os.Getwd()
	log.Println("current path:", wd)
	// https://stackoverflow.com/questions/12482702/pythons-os-chdir-and-os-getcwd-mismatch-when-using-tempfile-mkdtemp-on-ma

	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)

	// https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
	// As of Go 1.8 (Released February 2017) the recommended way of doing this is with os.Executable:
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
}

// examples:
// https://www.cnblogs.com/golove/p/5903579.html
func filepathPkg() {
	cpath, _ := os.Getwd()
	log.Println("current path:", cpath)
	tpath := filepath.Join(cpath, "../..///")

	fmt.Println(filepath.Clean(tpath))

	abs, _ := filepath.Abs(tpath)
	fmt.Println(abs)

	base := filepath.Base(tpath)
	fmt.Println(base)

	// seperator is ':'
	lpath := "a/b:cd"
	list := filepath.SplitList(lpath)
	fmt.Println(list)
}

func main() {
	execPath()
	filepathPkg()
}
