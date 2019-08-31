package main

import (
	"fmt"
	"net/http"

	"./config"
	"./routes"
)

func main() {
	router := routes.GetRouter()
	fmt.Println("listen on " + config.JSONConfig.Host + ":" + config.JSONConfig.Port)
	error := http.ListenAndServe("localhost:8585", router)
	if error != nil {
		fmt.Println(error.Error())
	}
}
