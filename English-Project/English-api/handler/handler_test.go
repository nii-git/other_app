package handler

import (
	"english-frequency/model"
	"english-frequency/usecase"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		logger    slog.Logger
		usecase   usecase.Usecase
		validator model.Validator
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.logger, tt.args.usecase, tt.args.validator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_HandlerFunc(t *testing.T) {
	type fields struct {
		logger    slog.Logger
		usecase   usecase.Usecase
		validator model.Validator
	}
	tests := []struct {
		name   string
		fields fields
		want   echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				logger:    tt.fields.logger,
				usecase:   tt.fields.usecase,
				validator: tt.fields.validator,
			}
			if got := h.HandlerFunc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.HandlerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_FrequencyHandlerFunc(t *testing.T) {
	type fields struct {
		logger    slog.Logger
		usecase   usecase.Usecase
		validator model.Validator
	}
	tests := []struct {
		name   string
		fields fields
		want   echo.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				logger:    tt.fields.logger,
				usecase:   tt.fields.usecase,
				validator: tt.fields.validator,
			}
			if got := h.FrequencyHandlerFunc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler.FrequencyHandlerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
