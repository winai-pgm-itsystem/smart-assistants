package databases

import (
	"encoding/json"
)

func UnmarshalExpenseResponse(data []byte) (ExpenseResponse, error) {
	var r ExpenseResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func UnmarshalAirtableErrorResponse(data []byte) (AirtableErrorResponse, error) {
	var r AirtableErrorResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type ExpenseResponse struct {
	Records []Record `json:"records"`
}

type Record struct {
	ID          string `json:"id"`
	CreatedTime string `json:"createdTime"`
	Fields      Fields `json:"fields"`
}

type Fields struct {
	Date     string  `json:"Date"`
	Category string  `json:"Category"`
	Amount   float32 `json:"Amount"`
	Remark   string  `json:"Remark,omitempty"`
}

type AirtableErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
