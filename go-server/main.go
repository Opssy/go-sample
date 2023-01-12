package go_server

import (
	"fmt"
	"log"
	"net/http"
)

func go_server() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("./", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Start serving at port 8080\n")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
