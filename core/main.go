package main

import (
	"MagicPotal/server"
	"log"
)

func main() {

	port := "8081"
	r := server.Setup()
	log.Println("listen on: ", port)

	err := r.Run(":" +port) //(viper.GetString("bindAddr") + ":" + viper.GetString("bindPort"))
	if err != nil {
		log.Println("Error running server: ", err.Error())
	}
}
