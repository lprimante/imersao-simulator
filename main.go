package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/lprimante/imersao-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("oieeee", "readtest", producer)
	fmt.Println("publicado")

	for {
		_ = 1
	}
	// route := route2.Route{
	// 	ID:       "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringJson, _ := route.ExportJsonPositions()
	// fmt.Println(stringJson[0])
}
