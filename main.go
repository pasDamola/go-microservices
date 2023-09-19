package main

import (
	"fmt"
	"log"
	"net/http"
)




func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!!")
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
