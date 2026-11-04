package waf

import (
	"net/http"
)

func inspect(next http.Handler) http.Handler {
    return a handler that, for each request:
        look at the query string
        if it contains something suspicious (start with just a single quote):
            respond 403 and stop
        otherwise call next to continue to the proxy
}