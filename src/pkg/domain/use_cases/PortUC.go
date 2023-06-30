package use_cases

import (
	"90poe/src/pkg/domain"
	"90poe/src/pkg/domain/ports"
	"context"
	"fmt"
	"io"
	"strings"
)

type PortUC struct {
	parser ports.PortParser
}

func NewPortUC(parser ports.PortParser) *PortUC {
	return &PortUC{parser: parser}
}

func save(port domain.Port) {

}

func (puc *PortUC) ParseAndPersist(ctx context.Context, reader io.Reader) error {
	portChannel := make(chan domain.Port)
	errChannel := make(chan error)

	defer func() {
		close(portChannel)
		close(errChannel)
	}()

	go puc.parser.ParserReader(reader, portChannel, errChannel)
	var i int64
	for {
		select {
		case port := <-portChannel:
			if strings.TrimSpace(port.Key) == "" {
				fmt.Printf("\nPortSaved: %d\n", i)
				return nil
			}

			i++
			save(port)
		case err := <-errChannel:
			return err
		}
	}
}
