package gateway

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type NotificationMailer struct{}

func (*NotificationMailer) SentMailer(mailer string) (err error) {
	url := "https://run.mocky.io/v3/0f35cc7d-833a-410c-ba07-7a0500fb9a2b"

	payload := map[string]interface{}{
		"email": mailer,
		"type":  "notification:transaction",
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return
	}

	req, _ := http.NewRequest("POST", url, &buf)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "picpay-simplificado")

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("SEND-MAILER-API-%d", res.StatusCode)
		return errors.New(msg)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(res)
	fmt.Println(string(body))
	return
}
