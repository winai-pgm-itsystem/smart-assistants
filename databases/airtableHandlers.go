package databases

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"smart-assistants/utils"
)

func AddExpenseHandler(category string, amount float32, remark string) (string, error) {
	date := utils.GetCurrentDate()
	requestBody := map[string]interface{}{
		"records": []map[string]interface{}{
			{
				"fields": map[string]interface{}{
					"Date":     date,
					"Category": category,
					"Amount":   amount,
					"Remark":   remark,
				},
			},
		},
	}

	resp, err := requestAirtable("POST", airtableURL(), requestBody)
	if err != nil {
		return "", fmt.Errorf("requst is invalid")
	}
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read res body: %s", err)
		}
		errRes, err := UnmarshalAirtableErrorResponse(bodyBytes)
		return "", fmt.Errorf("Airtable request failed message: %s code: %d", errRes.Error.Message, resp.StatusCode)
	}
	bodyBytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return "", fmt.Errorf("failed to read res body: %s", err)
	}
	result, err := UnmarshalExpenseResponse(bodyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %s", err)
	}
	if len(result.Records) > 0 {
		return result.Records[0].ID, nil
	}
	return "", fmt.Errorf("No records returned in the response")
}

func GetExpenseHandler(dateFilter, dateFotmat string) (*ExpenseResponse, error) {
	filterFormula := fmt.Sprintf("DATETIME_FORMAT({Date},'%s')='%s'", dateFotmat, dateFilter)
	sortField := "Date"
	sortDirection := "asc"
	query := fmt.Sprintf("?filterByFormula=%s&sort[0][field]=%s&sort[0][direction]=%s",
		filterFormula, sortField, sortDirection)
	resp, err := requestAirtable("GET", constructAirtableURL(query), map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("response error : %s", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read res body: %s", err)
		}
		errRes, err := UnmarshalAirtableErrorResponse(bodyBytes)
		return nil, fmt.Errorf("Airtable request failed message: %s code: %d", errRes.Error.Message, resp.StatusCode)
	}
	bodyBytes, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, fmt.Errorf("failed to read res body: %s", err)
	}
	result, err := UnmarshalExpenseResponse(bodyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %s", err)
	}

	return &result, nil
}
