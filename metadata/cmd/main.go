package cmd

import (
	"log"
	"net/http"

	"movie-service/metadata/internal/controller/metadata"
	httphandler "movie-service/metadata/internal/handler/http"
	"movie-service/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.New()
	ctrl := metadata.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
