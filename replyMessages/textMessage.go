package replymessages

import (
	"bytes"
	"fmt"
	"net/http"
)

func TextMessage(replyToken, message string) {
	url := "https://api.line.me/v2/bot/message/reply"

	requestBody := []byte(fmt.Sprintf(`{
    "replyToken": "%s",
    "messages": [
      {
        "type": "text",
        "text": "%s"
      }
    ]
  }`, replyToken, message))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer {RhS6tg3pgrEaImRK0uMqSriQ1sA+MSLdsnx+YctlkM4MkJAQlwKPD/7QcGx2RMxHDGf51d3BOxf10xGa14gcpjaToyZNDLJw7eNn8NwsBtfDEsFV9VFOPwkrh3/TpElZpZjpja4kajfZmCXUe/u8nQdB04t89/1O/w1cDnyilFU=}")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
