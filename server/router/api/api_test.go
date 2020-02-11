package api

import (
	"testing"

	"github.com/labstack/echo"
)

func Test_addAPI(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := addAPI(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("addAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
