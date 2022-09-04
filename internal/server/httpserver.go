package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/diyliv/keyvaluestorage/internal/models"
	services "github.com/diyliv/keyvaluestorage/internal/services/http"
)

type server struct {
	storage models.Storage
}

func NewServer(storage models.Storage) *server {
	return &server{storage: storage}
}

func (s *server) RunHttp() {
	services := services.NewHttpHandler(s.storage)
	log.Printf("Starting Http server on port :8080")

	http.HandleFunc("/api/add", services.Add)
	http.HandleFunc("/api/get", services.Get)
	http.HandleFunc("/api/delete", services.Delete)

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	<-done
}
