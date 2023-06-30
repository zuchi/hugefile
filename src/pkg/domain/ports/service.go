package ports

import (
	"90poe/src/pkg/domain"
	"context"
	"errors"
)

type Service struct {
	portRepository PortRepository
}

func NewServicePort(portRepository PortRepository) *Service {
	return &Service{portRepository: portRepository}
}

func (sp *Service) SavePort(ctx context.Context, port domain.Port) error {
	if port.Key == "" {
		return errors.New("there is no port to be saved")
	}

	sp.portRepository.Save(ctx, port)

	return nil
}

func (sp *Service) GetPortByKey(ctx context.Context, key string) (domain.Port, error) {

	return domain.Port{}, nil
}
