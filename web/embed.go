package web

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
	"strings"
)

//go:generate pnpm i
//go:generate pnpm build
//go:embed all:build
var webBuidlFs embed.FS

var webBuildDir, _ = fs.Sub(webBuidlFs, "build")

func EmbedWeb(mux *http.ServeMux) {
	slog.Info("Embedded web")

	fileServer := http.FileServerFS(webBuildDir)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			fsPath := strings.TrimPrefix(r.URL.Path, "/")

			_, err := fs.Stat(webBuildDir, fsPath)
			if err != nil {
				r.URL.Path = "/"
			}
		}

		fileServer.ServeHTTP(w, r)
	})
}
