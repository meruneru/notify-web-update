package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Scraping() (bool, error) {
	// 対象のウェブページ
	res, err := http.Get("https://eplus.jp/sf/detail/0158310001-P0030052?P6=001&P1=0402&P59=1")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return false, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// コンテンツ取得
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	// 指定の要素を取得
	statusElement := doc.Find(".ticket-status__item").First()
	if statusElement.Length() == 0 {
		log.Fatalf("HTML Element not found")
		return false, fmt.Errorf("HTML Element not found")
	}
	if statusElement.Text() != "予定枚数終了" {
		return true, nil
	}
	return false, nil
}
