package main

import (
	"github.com/diyliv/keyvaluestorage/internal/models"
	"github.com/diyliv/keyvaluestorage/internal/server"
	"github.com/diyliv/keyvaluestorage/pkg/realize"
)

func main() {
	keyval := realize.NewStorage(make(map[interface{}]models.Response))
	server := server.NewServer(keyval)
	server.RunHttp()

}
