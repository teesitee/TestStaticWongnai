package router

import (
	"reflect"
	"testing"
	"teststaticwongnai/api"

	"github.com/gin-gonic/gin"
)

func TestNewRouter(t *testing.T) {

	type args struct {
		h api.Handler
	}
	tests := []struct {
		name string
		args args
		want *gin.Engine
	}{
		{"case", args{}, gin.Default()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
