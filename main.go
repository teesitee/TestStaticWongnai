package main

import (
	"fmt"
	"teststaticwongnai/api"

	"teststaticwongnai/response"
	"teststaticwongnai/router"
)

func main() {
	fmt.Println("test")
	var r response.Respon

	h := api.Handler{
		ServiceRes: r,
	}
	fmt.Println("Start !!!")
	rou := router.NewRouter(h)
	rou.Run(":8080")

}
