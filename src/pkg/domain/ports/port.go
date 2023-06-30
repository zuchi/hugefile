package ports

import (
	"90poe/src/pkg/domain"
	"io"
)

type PortParser interface {
	ParserReader(reader io.Reader, nextPort chan domain.Port, errChannel chan error) error
}
