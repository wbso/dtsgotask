package frontend

// Embed compiled JS in go binary file
import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:build
var Static embed.FS

// Dist embed compiled JS in go binary file
func Dist() http.FileSystem {
	fsys, err := fs.Sub(Static, "build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
