package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/meruneru/notify-web-update/modules/line"
	"github.com/meruneru/notify-web-update/modules/web"
)

func main() {
	lambda.Start(Handler)
}

func Handler() (string, error) {
	isChanged, err := web.Scraping()
	if err != nil {
		log.Fatal("Error in Scraping: ", err)
		return "", err
	}

	if isChanged {
		err = line.SendLineNotify("The item is available now!")
	} else {
		err = line.SendLineNotify("The item is nothing")
	}
	return "ok", err
}
