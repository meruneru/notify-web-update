package line

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func SendLineNotify(message string) error {
	token := os.Getenv("LINE_NOTIFY_TOKEN")

	if token == "" {
		fmt.Println("LINE_NOTIFY_TOKEN is not set")
		os.Exit(1)
	}

	url := "https://notify-api.line.me/api/notify"
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte("message="+message)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal("Error in sending Line Notify: ", err)
	}

	return err
}
