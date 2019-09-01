package main

import (
	"fmt"
	"net/http"

	"lingotalk-exam/config"
	"lingotalk-exam/routes"
)

func main() {
	router := routes.GetRouter()
	fmt.Println("listen on " + config.JSONConfig.Host + ":" + config.JSONConfig.Port)
	error := http.ListenAndServe(fmt.Sprintf("%s:%s", config.JSONConfig.Host, config.JSONConfig.Port), router)
	if error != nil {
		fmt.Println(error.Error())
	}
}
