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

func SendFlexExpenseMessage(amount, category, currentToday, limit, day string) {
	url := "https://api.line.me/v2/bot/message/push"
	requestBody := []byte(getFlexExpenseStyle(os.Getenv("USER_ID"), amount, category, currentToday, limit, day))
	err := sendLineMessageApi(url, requestBody)
	if err != nil {
		fmt.Println("send flex message error:", err)
	}

}

func SendFlexPPQRMessage(amount string, total float64) {
	url := "https://api.line.me/v2/bot/message/push"
	imageUrl := fmt.Sprintf(`https://promptpay.io/%s/%v.png`, os.Getenv("PROMPT_PAY_ID"), total)
	requestBody := getFlexPPQRStyle(os.Getenv("USER_ID"), imageUrl, amount, os.Getenv("BANK_NAME"))
	err := sendLineMessageApi(url, []byte(requestBody))
	if err != nil {
		fmt.Println("send flex message error:", err)
	}

}
