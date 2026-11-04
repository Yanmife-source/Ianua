package waf 

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"log"

)


func ReverseProxy (req *http.Request) *http.Request{

	target,err:=url.Parse("http://localhost:8001")
	if err!=nil {
		log.Fatalf("Failed to parse target URL: %w", err)
	}
	proxy := &httputil.ReverseProxy{
		Rewrite: func(preq *httputil.ProxyRequest) {
			// See also ProxyRequest.SetURL.
			preq.SetURL(target)
			// This header cannot be affected by a malicious client.
			//preq.Out.Header.Set("Some-Header", "some-header-value")

			// X-Forwarded- headers sent by the client have been
			// removed from preq.Out.
			// ProxyRequest.SetXForwarded optionally adds new ones.
			preq.SetXForwarded()
		},
	}
}