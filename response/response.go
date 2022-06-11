package response

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"teststaticwongnai/duplicate"
)

type Respon struct{}

func (Respon) ResGet() (map[string]int, duplicate.Age) {
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
	return Dup_map, Dup_Age
}
