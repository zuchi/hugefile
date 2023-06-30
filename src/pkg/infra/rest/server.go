package rest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

const delim = 123

type Port struct {
	Key         string
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func CreateNewRestServer() {
	mux := http.NewServeMux()

	// upload file
	mux.HandleFunc("/upload", func(response http.ResponseWriter, request *http.Request) {
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

			dec := json.NewDecoder(reader)

			token, err := dec.Token()
			if err != nil {
				fmt.Printf("%v", err)
				writeMessage(response, "something went wrong", http.StatusInternalServerError)
				return
			}

			if d, ok := token.(json.Delim); !ok || d != json.Delim(delim) {
				writeMessage(response, "invalid token", http.StatusBadRequest)
				return
			}

			var i int64
			for dec.More() {
				p, err := parseObj(dec)
				if err != nil {
					writeMessage(response, "invalid token", http.StatusBadRequest)
					return
				}
				i++
				//fmt.Printf("%+v\n", p)
				save(p)
			}
			fmt.Printf("\nTotal: %d\n", i)

			return
		}

		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("this method is not supported"))
	})

	fmt.Printf("server is running at 3000\n")
	http.ListenAndServe(":3000", mux)
}

func parseObj(dec *json.Decoder) (Port, error) {
	var p Port

	obj, err := dec.Token()
	if err != nil {
		return Port{}, err
	}
	p.Key = obj.(string)

	err = dec.Decode(&p)
	if err != nil {
		return Port{}, err
	}
	return p, nil
}

func save(p Port) {

}

func writeMessage(response http.ResponseWriter, message string, statusCode int) {
	response.WriteHeader(statusCode)
	response.Write([]byte(message))
}
