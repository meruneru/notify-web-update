package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/meruneru/notify-web-update/notifier"
	"github.com/meruneru/notify-web-update/scraper"
)

func main() {
	lambda.Start(Handler)
}

func Handler() (string, error) {
	isChanged, err := scraper.Scraping()
	if err != nil {
		log.Fatal("Error in Scraping: ", err)
	}

	if isChanged {
		err = notifier.SendLineNotify("The item is available now!")
		if err != nil {
			log.Fatal("Error in sending Line Notify: ", err)
		}
	}
}
