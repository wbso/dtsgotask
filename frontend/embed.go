package frontend

// Embed compiled JS in go binary file
import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:build2
var Static embed.FS

// Dist embed compiled JS in go binary file
func Dist() http.FileSystem {
	fsys, err := fs.Sub(Static, "build2")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
