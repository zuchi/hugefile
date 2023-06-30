package use_cases

import (
	"90poe/src/pkg/domain"
	"90poe/src/pkg/domain/ports"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
)

const validJson = `
{
    "PORT1": {
		"Key": "PORT1",
    	"name": "Ajman",
    	"city": "Ajman",
    	"country": "United Arab Emirates",
    	"alias": [],
    	"regions": [],
    	"coordinates": [
       		 55.5136433,
        	 25.4052165
    	],
    	"province": "Ajman",
    	"timezone": "Asia/Dubai",
    	"unlocs": [
       		 "AEAJM"
    	],
    	"code": "Port1 Code"
    },
    "PORT2": {
		"Key": "PORT2",
    	"name": "PORT2 Name",
    	"city": "Port2 City",
    	"country": "Port2 Country",
    	"alias": [],
    	"regions": [],
    	"coordinates": [
       		 55.5136433,
        	 25.4052165
    	],
    	"province": "Ajman",
    	"timezone": "Asia/Dubai",
    	"unlocs": [
       		 "AEAJM"
    	],
    	"code": "Port2 Code"
    }
}
`

const invalidJson = `
[
    "PORT1": {
		"Key": "PORT1",
    	"name": "Ajman",
    	"city": "Ajman",
    	"country": "United Arab Emirates",
    	"alias": [],
    	"regions": [],
    	"coordinates": [
       		 55.5136433,
        	 25.4052165
    	],
    	"province": "Ajman",
    	"timezone": "Asia/Dubai",
    	"unlocs": [
       		 "AEAJM"
    	],
    	"code": "Port1 Code"
    },
    "PORT2": {
		"Key": "PORT2",
    	"name": "PORT2 Name",
    	"city": "Port2 City",
    	"country": "Port2 Country",
    	"alias": [],
    	"regions": [],
    	"coordinates": [
       		 55.5136433,
        	 25.4052165
    	],
    	"province": "Ajman",
    	"timezone": "Asia/Dubai",
    	"unlocs": [
       		 "AEAJM"
    	],
    	"code": "Port2 Code"
    },
]
`

func TestPortUC_ParseAndPersist(t *testing.T) {
	ctx := context.Background()

	type fields struct {
		parser      ports.PortParser
		portService ports.PortService
	}
	type args struct {
		ctx    context.Context
		reader io.Reader
	}

	type testStruct struct {
		name                        string
		fields                      fields
		args                        args
		numberExecutionSaveFunction int
		wantErr                     bool
	}

	tests := []testStruct{
		{
			name: "should execute correctly",
			fields: fields{
				parser: &ports.PortParserMock{ParserReaderFunc: func(reader io.Reader, nextPort chan domain.Port, errChannel chan error) {
					bytes, err := io.ReadAll(reader)
					if err != nil {
						errChannel <- err
					}

					objects := make(map[string]domain.Port)
					err = json.Unmarshal(bytes, &objects)
					if err != nil {
						errChannel <- err
					}

					for k, v := range objects {
						v.Key = k
						nextPort <- v
					}

					nextPort <- domain.Port{Key: ""} //finishing processing
				}},
				portService: &ports.PortServiceMock{SavePortFunc: func(ctx context.Context, port domain.Port) error {
					return nil
				}},
			},
			args: args{
				ctx:    ctx,
				reader: strings.NewReader(validJson),
			},
			numberExecutionSaveFunction: 2,
			wantErr:                     false,
		},
		{
			name: "should get an error because json is invalid",
			fields: fields{
				parser: &ports.PortParserMock{ParserReaderFunc: func(reader io.Reader, nextPort chan domain.Port, errChannel chan error) {
					bytes, err := io.ReadAll(reader)
					if err != nil {
						errChannel <- err
						return
					}

					objects := make(map[string]domain.Port)
					err = json.Unmarshal(bytes, &objects)
					if err != nil {
						errChannel <- err
						return
					}

					for k, v := range objects {
						v.Key = k
						nextPort <- v
					}

					nextPort <- domain.Port{Key: ""} //finishing processing
				}},
				portService: &ports.PortServiceMock{SavePortFunc: func(ctx context.Context, port domain.Port) error {
					return nil
				}},
			},
			args: args{
				ctx:    ctx,
				reader: strings.NewReader(invalidJson),
			},
			numberExecutionSaveFunction: 0,
			wantErr:                     true,
		},
		{
			name: "should get error because save function has returned error",
			fields: fields{
				parser: &ports.PortParserMock{ParserReaderFunc: func(reader io.Reader, nextPort chan domain.Port, errChannel chan error) {
					nextPort <- domain.Port{Key: "123"}
				}},
				portService: &ports.PortServiceMock{SavePortFunc: func(ctx context.Context, port domain.Port) error {
					return fmt.Errorf("something wrong happened when saving")
				}},
			},
			args: args{
				ctx:    ctx,
				reader: strings.NewReader(validJson),
			},
			numberExecutionSaveFunction: 1,
			wantErr:                     true,
		},
	}

	for _, tt := range tests {
		puc := NewPortUC(tt.fields.parser, tt.fields.portService)
		t.Run(tt.name, func(t *testing.T) {
			if err := puc.ParseAndPersist(tt.args.ctx, tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("ParseAndPersist() error = %v, wantErr %v", err, tt.wantErr)
			}

			portServiceMock := tt.fields.portService.(*ports.PortServiceMock)
			if len(portServiceMock.SavePortCalls()) != tt.numberExecutionSaveFunction {
				t.Errorf("expected to be executed SavePort function: %d but executed: %d", tt.numberExecutionSaveFunction, len(portServiceMock.SavePortCalls()))
			}
		})
	}
}
func x() {

}
