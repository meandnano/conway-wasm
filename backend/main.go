package main

import (
	"net/http"

	"github.com/meandnano/conway-web/backend/assets"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, assets.FS, "index.html")
	})

	http.Handle("GET /assets/*", http.StripPrefix("/assets/", http.FileServerFS(assets.FS)))

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		panic(err)
	}
}
