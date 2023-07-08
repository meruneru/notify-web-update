package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler() {
	isChanged, err := Scraping()
	if err != nil {
		log.Fatal("Error in Scraping: ", err)
	}

	if isChanged {
		err = SendLineNotify("The item is available now!")
		if err != nil {
			log.Fatal("Error in sending Line Notify: ", err)
		}
	}
}
