package ports

import (
	"90poe/src/pkg/domain"
	"context"
	"errors"
	"strings"
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

	err := sp.portRepository.Save(ctx, port)
	if err != nil {
		return err
	}

	return nil
}

func (sp *Service) GetPortByKey(ctx context.Context, key string) (domain.Port, error) {
	if strings.TrimSpace(key) == "" {
		return domain.Port{}, nil
	}
	return sp.portRepository.FindByKey(ctx, key)
}
