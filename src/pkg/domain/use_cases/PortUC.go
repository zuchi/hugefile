package use_cases

import (
	"90poe/src/pkg/domain"
	"90poe/src/pkg/domain/ports"
	"context"
	"go.uber.org/zap"
	"io"
	"strings"
)

type PortUC struct {
	parser      ports.PortParser
	portService ports.PortService
	log         *zap.SugaredLogger
}

func NewPortUC(parser ports.PortParser, service ports.PortService) *PortUC {

	logger, _ := zap.NewProduction()
	log := logger.Sugar().With("component", "port uc")

	return &PortUC{
		parser:      parser,
		portService: service,
		log:         log,
	}
}

func (puc *PortUC) ParseAndPersist(ctx context.Context, reader io.Reader) error {
	portChannel := make(chan domain.Port)
	errChannel := make(chan error)

	defer func() {
		close(portChannel)
		close(errChannel)
	}()

	var i int64

	go puc.parser.ParserReader(reader, portChannel, errChannel)

	for {
		select {
		case port := <-portChannel:
			// if we receive an empty key, it means that we don't have more things to process
			if strings.TrimSpace(port.Key) == "" {
				puc.log.Infof("total processed: %d", i)
				return nil
			}

			i++
			if i%250 == 0 {
				puc.log.Infof("already processed: %d", i)
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
