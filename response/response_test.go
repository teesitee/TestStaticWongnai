package response

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"testing"
	"teststaticwongnai/duplicate"
)

func TestRespon_ResGet(t *testing.T) {
	resp, err := http.Get("http://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		fmt.Println(err)

	}

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err2 := json.Unmarshal(bodyByte, &duplicate.In)
	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}

	Dup_map := (duplicate.In).Dupcount()
	Dup_Age := (duplicate.In).Agecount()

	tests := []struct {
		name  string
		r     Respon
		want  map[string]int
		want1 duplicate.Age
	}{
		{"case 1", Respon{}, Dup_map, Dup_Age},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.r.ResGet()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Respon.ResGet() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Respon.ResGet() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
