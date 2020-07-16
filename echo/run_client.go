package main

import (
	"echo/impl/v1/client"
	"log"
)

func main() {
	log.Println("[.] Runing Client")
	client.Run()
	log.Println("[o] Runing Client")
}
