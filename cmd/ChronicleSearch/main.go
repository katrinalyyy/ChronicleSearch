package main

import (
	"Lab1/intermal/api"
	"log"
)

func main() {
	log.Println("Application start!")
	api.StartServer()
	log.Println("Application terminated")
}
