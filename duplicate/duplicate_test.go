package duplicate

import (
	"reflect"
	"testing"
)

func TestResponse_Dupcount(t *testing.T) {
	tests := []struct {
		name  string
		listP Response
		want  map[string]int
	}{
		{"case1", Response{}, map[string]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.listP.Dupcount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Dupcount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Agecount(t *testing.T) {
	tests := []struct {
		name  string
		listA Response
		want  Age
	}{
		{"case2", Response{}, Age{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.listA.Agecount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.Agecount() = %v, want %v", got, tt.want)
			}
		})
	}
}
