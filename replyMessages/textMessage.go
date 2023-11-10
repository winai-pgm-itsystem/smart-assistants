package replymessages

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer {%s}", os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
