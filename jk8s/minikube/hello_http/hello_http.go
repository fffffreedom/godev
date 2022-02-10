package main

import (
	"net/http"
	"fmt"
)

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello http")
}

func main ()  {
	http.HandleFunc("/", HelloHandle)
	http.ListenAndServe(":8080", nil)
}
