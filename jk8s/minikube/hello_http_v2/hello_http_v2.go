package main

import (
	"net/http"
	"fmt"
)

func HelloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello http v2")
}

func main ()  {
	http.HandleFunc("/", HelloHandle)
	http.ListenAndServe(":8080", nil)
}
