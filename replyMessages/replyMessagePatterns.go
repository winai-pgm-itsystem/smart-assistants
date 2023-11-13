package replymessages

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func sendLineMessageApi(url string, requestBody []byte) error {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer {%s}", os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func SendTextMessage(replyToken, message string) {
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

	err := sendLineMessageApi(url, requestBody)
	if err != nil {
		fmt.Println("send text message error:", err)
	}

}

func SendFlexMessage(amount, category, currentToday, limit, day string) {
	url := "https://api.line.me/v2/bot/message/push"

	requestBody := []byte(getFlexExpenseStyle(os.Getenv("USER_ID"), amount, category, currentToday, limit, day))

	err := sendLineMessageApi(url, requestBody)
	if err != nil {
		fmt.Println("send flex message error:", err)
	}

}
