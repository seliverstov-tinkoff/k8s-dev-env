package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	foo := os.Getenv("FOO")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(foo))
		fmt.Println("request test")
	})
	http.ListenAndServe(":8080", nil)
}
