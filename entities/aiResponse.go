package entities

import "encoding/json"

func UnmarshalAIResponse(data []byte) (AIResponse, error) {
	var r AIResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AIResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AIResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Output        string         `json:"output"`
	SafetyRatings []SafetyRating `json:"safetyRatings"`
}

type SafetyRating struct {
	Category    string `json:"category"`
	Probability string `json:"probability"`
}
