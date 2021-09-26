// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// body
func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)) // 打出body

	// stream 只能读一次 单向流
	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)) // 什么也不打出
}

// getbody
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)) // 打出body

	// stream 只能读一次 单向流
	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body)) // 什么也不打出
}

func main() {
	http.HandleFunc("/user", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
