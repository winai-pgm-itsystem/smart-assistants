package replymessages

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"smart-assistants/entities"
)

func aiGenearateText(message string) (string, error) {

	host := "https://generativelanguage.googleapis.com/v1beta3/models/text-bison-001:generateText"
	url := fmt.Sprintf(`%s?key=%s`, host, os.Getenv("PALM_KEY"))

	requestBody := []byte(fmt.Sprintf(`{
		"prompt": {
			"text": "%s"
		}
	}`, message))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read res body: %s", err)
		}

		return "", fmt.Errorf("ai  request failed message: %s", string(bodyBytes))
	}
	bodyBytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return "", fmt.Errorf("failed to read res body: %s", err)
	}
	result, err := entities.UnmarshalAIResponse(bodyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %s", err)
	}

	return result.Candidates[0].Output, nil

}
