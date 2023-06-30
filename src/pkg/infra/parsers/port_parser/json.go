package port_parser

import (
	"90poe/src/pkg/domain"
	"encoding/json"
	"fmt"
	"io"
)

const delim = 123 // this is the { rune character

type Json struct {
}

func NewJsonParser() *Json {
	return &Json{}
}

func (j *Json) ParserReader(reader io.Reader, nextPort chan domain.Port, errChannel chan error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("channels has been closed")
		}
	}()

	dec := json.NewDecoder(reader)
	token, err := dec.Token()
	if err != nil {
		errChannel <- err
		return
	}

	if d, ok := token.(json.Delim); !ok || d != json.Delim(delim) {
		errChannel <- err
		return
	}

	for dec.More() {
		var p domain.Port
		obj, err := dec.Token()
		if err != nil {
			errChannel <- err
			return
		}

		p.Key = obj.(string)

		err = dec.Decode(&p)
		if err != nil {
			errChannel <- err
			return
		}

		nextPort <- p
	}
	nextPort <- domain.Port{Key: ""}
	return
}
