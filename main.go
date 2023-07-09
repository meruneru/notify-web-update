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
		log.Printf("変更を検出")
		err = line.SendLineNotify("商品が販売されています！！！ https://eplus.jp/sf/detail/0158310001-P0030052?P6=001&P1=0402&P59=1")
	} else {
		log.Printf("変更なし")
		//err = line.SendLineNotify("商品が売り切れです。")
	}
	return "ok", err
}
