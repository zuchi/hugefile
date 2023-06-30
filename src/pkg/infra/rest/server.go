package rest

import (
	"90poe/src/pkg/domain/use_cases"
	"context"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	PortUC *use_cases.PortUC
	log    *zap.SugaredLogger
	mux    *http.Server
}

func NewServer(portUC *use_cases.PortUC) *Server {
	srv := new(Server)
	logger, _ := zap.NewProduction()
	logSuggar := logger.Sugar().With("component", "server")
	srv.log = logSuggar

	mux := http.NewServeMux()
	// upload file
	mux.HandleFunc("/upload", srv.UploadFilePOST())
	// get by key
	mux.HandleFunc("/port", srv.GetPortGET())

	server := &http.Server{
		Handler: mux,
	}

	srv.mux = server
	srv.PortUC = portUC

	return srv
}

func (s *Server) Run(address string) {
	s.log.Infof("Starting server at %s", address)
	http.ListenAndServe(address, s.mux.Handler)
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.mux.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func writeMessage(response http.ResponseWriter, message string, statusCode int) {
	response.WriteHeader(statusCode)
	response.Write([]byte(message))
}
