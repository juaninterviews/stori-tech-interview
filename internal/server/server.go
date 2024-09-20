package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Port    int
	Handler IHandler
}

func NewServer(
	port int) *Server {
	return &Server{
		Port:    port,
		Handler: NewHandler(),
	}
}

func (s *Server) Init() {
	//	Migrate method
	http.HandleFunc("/v1/migrate", s.Handler.MigrateHandler)

	fmt.Println("Server started successfully http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server startup httpError:", err)
	}
}
