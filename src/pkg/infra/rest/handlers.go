package rest

import (
	"bufio"
	"fmt"
	"net/http"
)

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
				writeMessage(response, "something went wrong", http.StatusInternalServerError)
				return
			}
		}
	}
}
