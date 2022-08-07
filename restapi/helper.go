package restapi

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"

	"github.com/wbso/dtsgotask/frontend"
)

// wrap the response writer to record the status code
type responseWrapper struct {
	http.ResponseWriter
	status int
}

func (wrp *responseWrapper) WriteHeader(status int) {
	// dont write anything if the status is not OK
	if status != http.StatusOK {
		wrp.status = status
		return
	}

	wrp.ResponseWriter.WriteHeader(status)
}

func (wrp *responseWrapper) Write(b []byte) (int, error) {
	// dont write anything if the status is not OK
	if wrp.status != http.StatusOK {
		return len(b), nil
	}
	return wrp.ResponseWriter.Write(b)
}

// HandleEmbeddedIndexFile returns a handler that serves the frontend file from embedded filesystem.
// If the requested file is not found, it serves the index.html.
func HandleEmbeddedIndexFile() http.HandlerFunc {
	fileServer := http.FileServer(frontend.Dist())
	return func(w http.ResponseWriter, r *http.Request) {
		wrapper := &responseWrapper{w, http.StatusOK}
		// filePath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		filePath := path.Clean(r.URL.Path)
		r.URL.Path = filePath
		// server the file
		fileServer.ServeHTTP(wrapper, r)
		// If the requested file is not found, it serves the index.html.
		if wrapper.status != http.StatusOK {
			wrapper.status = http.StatusOK
			wrapper.Header().Set("Content-Type", "text/html")
			r.URL.Path = "/"
			fileServer.ServeHTTP(wrapper, r)
		}
	}
}

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
