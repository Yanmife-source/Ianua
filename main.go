package main

import (
	"fmt"
	//"log/slog"
	"net/http"
	"Ianua/internal"
)

func main(){
	proxy:=waf.ReverseProxy("http://127.0.0.1:8001")

	wafProxy:=waf.Inspect(proxy)
	
	// 2. Start the HTTP server listening on port 8080
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", wafProxy)
	if err != nil {
		panic(err)
	}
}