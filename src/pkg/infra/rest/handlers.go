package rest

import (
	"bufio"
	"encoding/json"
	"net/http"
)

func (s *Server) GetPortGET() func(response http.ResponseWriter, request *http.Request) {
	log := s.log.With("component", "GetPortGET function")
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			id := request.URL.Query().Get("id")
			if id == "" {
				log.Infof("cannot get id from request")
				writeMessage(response, "there is no param sent", http.StatusBadRequest)
				return
			}

			port, err := s.PortUC.GetPortByKey(request.Context(), id)
			if err != nil {
				log.Errorf("cannot perform GetPortByKey: %v", err)
				writeMessage(response, "cannot get port from database", http.StatusInternalServerError)
				return
			}
			jsonPort, err := json.Marshal(port)
			if err != nil {
				log.Errorf("cannot unmarshal port: %v", err)
				writeMessage(response, "cannot send the port object to client", http.StatusInternalServerError)
				return
			}

			response.Header().Add("content-type", "application/json")
			response.WriteHeader(http.StatusOK)
			response.Write(jsonPort)

			log.Infof("searching by id %s performed sucessfully", id)
		}
	}
}

func (s *Server) UploadFilePOST() func(response http.ResponseWriter, request *http.Request) {
	log := s.log.With("component", "UploadFilePOST function")
	return func(response http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {

			file, header, err := request.FormFile("file")
			if err != nil {
				log.Errorf("cannot get file from request: %v", err)
				writeMessage(response, "cannot open file", http.StatusInternalServerError)
				return
			}

			defer func() {
				err := file.Close()
				if err != nil {
					log.Errorf("cannot close file on request: %v", err)
				}
			}()

			reader := bufio.NewReader(file)

			err = s.PortUC.ParseAndPersist(request.Context(), reader)
			if err != nil {
				log.Errorf("cannot parse and persist file: %v", err)
				writeMessage(response, "something went wrong when processing the file", http.StatusInternalServerError)
				return
			}

			log.Infof("file processed with success: %s", header.Filename)
		}
	}
}
