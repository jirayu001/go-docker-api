package model

type DataUsers struct {
	LineToken    string `json:"lineToken"`
	DisplayName  string `json:"displayName"`
	Email        string `json:"email"`
	DisplayImage string `json:"displayImage"`
}
type Result struct {
	Status int `json:"status"`
	Data   struct {
		LiffURL string `json:"liffURL"`
	} `json:"data"`
	Message string `json:"message"`
}
