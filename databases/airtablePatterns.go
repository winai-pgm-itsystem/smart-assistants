package databases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// const (
// 	airtableURL = "https://api.airtable.com/v0/appcich0xHqADv8NE/ExpenseTracker"
// )

func buildFilterFormula(dateFilter, dateFormat string) string {
	return fmt.Sprintf("DATETIME_FORMAT({Date}, '%s')='%s'", dateFormat, dateFilter)
}

func airtableURL() string {
	return fmt.Sprintf(`https://api.airtable.com/v0/%s/ExpenseTracker`, os.Getenv("AIR_TABLE_APP_ID"))
}

func constructAirtableURL(query string) string {
	return airtableURL() + query
}

func requestAirtable(method, url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("json marshel error : %s", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("NewRequest error : %s", err)
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("AIR_TABLE_ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	return client.Do(req)
}
