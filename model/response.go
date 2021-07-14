package model

type DataUsers struct {
	Token string `json:"token" example:"test"`
}
type Result struct {
	Status int `json:"status" example:"400"`
	Data   struct {
		LiffURL string `json:"liffURL" example:"https://f0ac1522c75e.ngrok.io"`
	} `json:"data"`
	Message string `json:"message" example:"กรุณากรอกข้อมูลให้ครบ"`
}
