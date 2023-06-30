package ports

import (
	"90poe/src/pkg/domain"
	"context"
	"io"
)

//go:generate moq -out ports_mock.go . PortParser PortRepository PortService
type PortParser interface {
	ParserReader(reader io.Reader, nextPort chan domain.Port, errChannel chan error)
}

type PortRepository interface {
	Save(ctx context.Context, port domain.Port) error
	FindByKey(ctx context.Context, key string) (domain.Port, error)
}

type PortService interface {
	SavePort(ctx context.Context, port domain.Port) error
	GetPortByKey(ctx context.Context, key string) (domain.Port, error)
}
