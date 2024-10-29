package main


import (
	"net/http"
	"fmt"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Println("Server running on port 1234 ...")
	http.ListenAndServe("localhost: 1234", nil)
}