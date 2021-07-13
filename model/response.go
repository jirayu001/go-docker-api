package model

type DataUsers struct {
	Token string `json:"token" binding:"required`
}
type Result struct {
	Status int `json:"status"`
	Data   struct {
		LiffURL string `json:"liffURL"`
	} `json:"data"`
	Message string `json:"message"`
}
