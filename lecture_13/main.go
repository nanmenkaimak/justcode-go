package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/static/{file_name}", sendFiles)
	http.ListenAndServe(":8080", r)
}

func sendFiles(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "file_name")

	fileBytes, err := os.ReadFile(fmt.Sprintf("internal/files/%s", fileName))
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}
