package ports

import (
	"90poe/src/pkg/domain"
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestService_GetPortByKey(t *testing.T) {
	ctx := context.Background()
	portExample := domain.Port{Key: "fetch"}

	type fields struct {
		portRepository PortRepository
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name                       string
		fields                     fields
		args                       args
		want                       domain.Port
		wantErrMessage             string
		numberOfFindByKeyExecution int
	}{
		{
			name: "should fetch port by key",
			fields: fields{portRepository: &PortRepositoryMock{FindByKeyFunc: func(ctx context.Context, key string) (domain.Port, error) {
				if key == "fetch" {
					return portExample, nil
				}
				return domain.Port{}, nil
			}}},
			args: args{
				ctx: ctx,
				key: "fetch",
			},
			want:                       portExample,
			wantErrMessage:             "",
			numberOfFindByKeyExecution: 1,
		},
		{
			name: "should fetch an empty port because key doesn't exist",
			fields: fields{portRepository: &PortRepositoryMock{FindByKeyFunc: func(ctx context.Context, key string) (domain.Port, error) {
				if key == "fetch" {
					return portExample, nil
				}
				return domain.Port{}, nil
			}}},
			args: args{
				ctx: ctx,
				key: "nokey",
			},
			want:                       domain.Port{},
			wantErrMessage:             "",
			numberOfFindByKeyExecution: 1,
		},
		{
			name: "should fetch an empty port because key was not provided",
			fields: fields{portRepository: &PortRepositoryMock{FindByKeyFunc: func(ctx context.Context, key string) (domain.Port, error) {
				if key == "fetch" {
					return portExample, nil
				}
				return domain.Port{}, nil
			}}},
			args: args{
				ctx: ctx,
				key: "",
			},
			want:                       domain.Port{},
			wantErrMessage:             "",
			numberOfFindByKeyExecution: 0,
		},
		{
			name: "should get an error because repository has failed",
			fields: fields{portRepository: &PortRepositoryMock{FindByKeyFunc: func(ctx context.Context, key string) (domain.Port, error) {
				return domain.Port{}, fmt.Errorf("some error")
			}}},
			args: args{
				ctx: ctx,
				key: "somekey",
			},
			want:                       domain.Port{},
			wantErrMessage:             "some error",
			numberOfFindByKeyExecution: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := &Service{
				portRepository: tt.fields.portRepository,
			}
			got, err := sp.GetPortByKey(tt.args.ctx, tt.args.key)
			if (err != nil) != (len(tt.wantErrMessage) > 0) || (err != nil && err.Error() != tt.wantErrMessage) {
				t.Errorf("SavePort() error message = %v, wantErr message %v", err, tt.wantErrMessage)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPortByKey() got = %v, want %v", got, tt.want)
			}

			mockRepo := tt.fields.portRepository.(*PortRepositoryMock)
			if len(mockRepo.FindByKeyCalls()) != tt.numberOfFindByKeyExecution {
				t.Errorf("expected to have %d number of calls, but received %d", tt.numberOfFindByKeyExecution, len(mockRepo.FindByKeyCalls()))
			}
		})
	}
}

func TestService_SavePort(t *testing.T) {

	ctx := context.Background()

	type fields struct {
		portRepository PortRepository
	}
	type args struct {
		ctx  context.Context
		port domain.Port
	}
	tests := []struct {
		name                            string
		fields                          fields
		args                            args
		wantErrMessage                  string
		numberOfExecutionSaveRepository int
	}{
		{
			name: "should save port without error",
			fields: fields{portRepository: &PortRepositoryMock{SaveFunc: func(ctx context.Context, port domain.Port) error {
				return nil
			}}},
			args: args{
				ctx:  ctx,
				port: domain.Port{Key: "id"},
			},
			wantErrMessage:                  "",
			numberOfExecutionSaveRepository: 1,
		},
		{
			name: "should get error because we don't have key",
			fields: fields{portRepository: &PortRepositoryMock{SaveFunc: func(ctx context.Context, port domain.Port) error {
				return nil
			}}},
			args: args{
				ctx:  ctx,
				port: domain.Port{Key: ""},
			},
			wantErrMessage:                  "there is no port to be saved",
			numberOfExecutionSaveRepository: 0,
		},
		{
			name: "should get error because repository sent an error",
			fields: fields{portRepository: &PortRepositoryMock{SaveFunc: func(ctx context.Context, port domain.Port) error {
				return fmt.Errorf("could not save")
			}}},
			args: args{
				ctx:  ctx,
				port: domain.Port{Key: "123"},
			},
			wantErrMessage:                  "could not save",
			numberOfExecutionSaveRepository: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sp := &Service{
				portRepository: tt.fields.portRepository,
			}

			err := sp.SavePort(tt.args.ctx, tt.args.port)
			if (err != nil) != (len(tt.wantErrMessage) > 0) || (err != nil && err.Error() != tt.wantErrMessage) {
				t.Errorf("SavePort() error message = %v, wantErr message %v", err, tt.wantErrMessage)
			}

			mock := tt.fields.portRepository.(*PortRepositoryMock)
			if len(mock.SaveCalls()) != tt.numberOfExecutionSaveRepository {
				t.Errorf("expected to have %d number of calls, but received %d", tt.numberOfExecutionSaveRepository, len(mock.SaveCalls()))
			}
		})
	}
}
