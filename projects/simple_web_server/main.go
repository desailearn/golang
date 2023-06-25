package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(resp, "404 Invalid Path", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(resp, "Invalid Method", http.StatusNotFound)
		return
	}
	fmt.Fprintf(resp, "hello!")
}

func formHandler(resp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(resp, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(resp, "POST request successful")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(resp, "Name = %s\n", name)
	fmt.Fprintf(resp, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
