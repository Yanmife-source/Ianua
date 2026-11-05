package waf 

import (
	"net/http/httputil"
	"net/url"
	"log"

)


func ReverseProxy (targetURL string) *httputil.ReverseProxy{

	target,err:=url.Parse(targetURL)
	if err!=nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}
	return &httputil.ReverseProxy{
		Rewrite: func(preq *httputil.ProxyRequest) {
			// See also ProxyRequest.SetURL.
			preq.SetURL(target)
			preq.SetXForwarded()
		},
	}
	
}