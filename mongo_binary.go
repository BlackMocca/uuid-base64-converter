package ub64c

import "encoding/json"

type MongoBinary struct {
	Data MongoBinaryData `json:"$binary"`
}

type MongoBinaryData struct {
	Base64  string `json:"base64"`
	SubType string `json:"subType"`
}

func NewMongoBinary(base64 string) MongoBinary {
	return MongoBinary{
		Data: MongoBinaryData{
			Base64:  base64,
			SubType: "4",
		},
	}
}

func (m MongoBinary) ToJSON() string {
	bu, _ := json.Marshal(m)
	return string(bu)
}
