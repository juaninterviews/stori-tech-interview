package main

import (
	"github.com/juaninterviews/stori-tech-interview/internal/server"
)

func main() {
	httpServer := server.NewServer(8080)

	httpServer.Init()
}
