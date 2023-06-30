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
	parser      ports.PortParser
	portService *ports.Service
}

func NewPortUC(parser ports.PortParser, service *ports.Service) *PortUC {
	return &PortUC{
		parser:      parser,
		portService: service,
	}
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

			err := puc.portService.SavePort(ctx, port)
			if err != nil {
				return err
			}
		case err := <-errChannel:
			return err
		}
	}
}

func (puc *PortUC) GetPortByKey(ctx context.Context, id string) (domain.Port, error) {
	return puc.portService.GetPortByKey(ctx, id)
}
