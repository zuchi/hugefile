package rest

import (
	"90poe/src/pkg/domain/use_cases"
	"fmt"
	"net/http"
)

type Server struct {
	PortUC *use_cases.PortUC
}

func NewServer(portUC *use_cases.PortUC) *Server {
	return &Server{PortUC: portUC}
}

func (s *Server) StartServer(address string) {
	mux := http.NewServeMux()

	// upload file
	mux.HandleFunc("/upload", s.UploadFilePOST())

	fmt.Printf("server is running at 3000\n")
	http.ListenAndServe(address, mux)
}

func writeMessage(response http.ResponseWriter, message string, statusCode int) {
	response.WriteHeader(statusCode)
	response.Write([]byte(message))
}
