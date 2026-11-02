package main

import (
	"fmt"
	//"log/slog"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})
	fmt.Println("Starting server on :8000...")
	
	// 2. Start the HTTP server listening on port 8000
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}