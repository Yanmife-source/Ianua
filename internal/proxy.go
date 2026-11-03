package waf 

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func ReverseProxy (req *http.Request) *http.Request{
	proxy := &httputil.ReverseProxy{
		Rewrite: func(preq *httputil.ProxyRequest) {
			// See also ProxyRequest.SetURL.
			preq.Out.URL.Scheme = "https"
			preq.Out.URL.Host = proxyHost

			// This header cannot be affected by a malicious client.
			preq.Out.Header.Set("Some-Header", "some-header-value")

			// X-Forwarded- headers sent by the client have been
			// removed from preq.Out.
			// ProxyRequest.SetXForwarded optionally adds new ones.
			preq.SetXForwarded()
		},
	}
}