package rest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) GetPortGET() func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			id := request.URL.Query().Get("id")
			if id == "" {
				writeMessage(response, "there is no param sent", http.StatusBadRequest)
				return
			}

			port, err := s.PortUC.GetPortByKey(request.Context(), id)
			if err != nil {
				writeMessage(response, "cannot get port from database", http.StatusInternalServerError)
				return
			}
			jsonPort, err := json.Marshal(port)
			if err != nil {
				writeMessage(response, "cannot send the port object to client", http.StatusInternalServerError)
				return
			}

			response.Header().Add("content-type", "application/json")
			response.WriteHeader(http.StatusOK)
			response.Write(jsonPort)
		}
	}
}

func (s *Server) UploadFilePOST() func(response http.ResponseWriter, request *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {

			file, _, err := request.FormFile("file")
			if err != nil {
				writeMessage(response, "cannot open file", http.StatusInternalServerError)
				return
			}

			defer func() {
				err := file.Close()
				if err != nil {
					fmt.Printf("%v", err)
				}
			}()

			reader := bufio.NewReader(file)

			err = s.PortUC.ParseAndPersist(request.Context(), reader)
			if err != nil {
				fmt.Printf("PArseAndPersist Error: %w\n", err)
				writeMessage(response, "something went wrong", http.StatusInternalServerError)
				return
			}
		}
	}
}
