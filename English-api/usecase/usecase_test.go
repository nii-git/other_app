package usecase

import (
	"english-frequency/infra"
	"english-frequency/model"
	"reflect"
	"testing"

	"golang.org/x/exp/slog"
)

func TestNewUsecase(t *testing.T) {
	type args struct {
		logger slog.Logger
		db     infra.DB
	}
	tests := []struct {
		name string
		args args
		want *Usecase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsecase(tt.args.logger, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_FrequencyUsecase(t *testing.T) {
	type fields struct {
		logger slog.Logger
		db     infra.DB
	}
	type args struct {
		request model.Frequency_request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Frequency_response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Usecase{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			got, err := u.FrequencyUsecase(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.FrequencyUsecase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.FrequencyUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
