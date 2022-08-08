package restapi

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxy(remoteURL string) http.HandlerFunc {
	remote, err := url.Parse(remoteURL)
	if err != nil {
		panic("failed parse url")
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	return func(w http.ResponseWriter, r *http.Request) {
		// r.Header.Add("X-Forwarded-Host", r.Host)
		// r.Header.Add("X-Forwarded-For", r.Host)
		// r.Header.Add("X-Forwarded-Proto", "https")
		// r.Header.Add("X-Origin-Host", r.Host)
		r.URL.Scheme = "https"
		r.URL.Host = remote.Host
		r.Host = remote.Host
		proxy.ServeHTTP(w, r)
	}
}
