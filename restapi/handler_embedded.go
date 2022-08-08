package restapi

import (
	"net/http"
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
		// filePath := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		filePath := path.Clean(r.URL.Path)
		r.URL.Path = filePath

		// serve the file, and write to the responseWrapper, so we can record the status code
		writeWrapper := &responseWrapper{w, http.StatusOK}
		fileServer.ServeHTTP(writeWrapper, r)

		// If the requested file is not found, it serve the index.html.
		if writeWrapper.status != http.StatusOK {
			writeWrapper.status = http.StatusOK
			writeWrapper.Header().Set("Content-Type", "text/html")
			r.URL.Path = "/"
			fileServer.ServeHTTP(writeWrapper, r)
		}
	}
}
