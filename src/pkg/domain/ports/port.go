package ports

import (
	"90poe/src/pkg/domain"
	"context"
	"io"
)

type PortParser interface {
	ParserReader(reader io.Reader, nextPort chan domain.Port, errChannel chan error) error
}

type PortRepository interface {
	Save(ctx context.Context, port domain.Port) error
	FindByKey(ctx context.Context, key string) (domain.Port, error)
}
