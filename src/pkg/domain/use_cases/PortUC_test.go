package use_cases

import (
	"90poe/src/pkg/domain/ports"
	"context"
	"io"
	"testing"
)

func TestPortUC_ParseAndPersist(t *testing.T) {
	type fields struct {
		parser      ports.PortParser
		portService *ports.Service
		log         *zap.SugaredLogger
	}
	type args struct {
		ctx    context.Context
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			puc := &PortUC{
				parser:      tt.fields.parser,
				portService: tt.fields.portService,
				log:         tt.fields.log,
			}
			if err := puc.ParseAndPersist(tt.args.ctx, tt.args.reader); (err != nil) != tt.wantErr {
				t.Errorf("ParseAndPersist() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
