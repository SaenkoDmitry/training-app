package web

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed dist/*
var content embed.FS

func SPAHandler() http.Handler {
	sub, _ := fs.Sub(content, "dist")
	fsHandler := http.FileServer(http.FS(sub))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// 1. API пути обрабатываются отдельно
		if strings.HasPrefix(path, "/api/") {
			http.NotFound(w, r)
			return
		}

		// 2. Проверяем, существует ли файл
		if _, err := sub.Open(strings.TrimPrefix(path, "/")); err == nil {
			fsHandler.ServeHTTP(w, r)
			return
		}

		// 3. Если файл не найден, отдаем index.html для React Router
		r.URL.Path = "/index.html"
		fsHandler.ServeHTTP(w, r)
	})
}
