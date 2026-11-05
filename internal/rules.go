package waf

import (
	"net/http"
	"strings"
)

func Inspect(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// WAF LOGIC HERE: Inspect request
		// Example: Block a specific User-Agent or malicious path
		for _,values:=range r.URL.Query(){
			for _,value:=range values{
				if strings.Contains(value,"'"){
					http.Error(w,"Forbidden by Web Applcation Firewall",http.StatusForbidden)
					return
				}
			}
		}

		// If the request is safe, pass it to the Reverse Proxy
		next.ServeHTTP(w, r)
	})
}