package signup

import (
	"encoding/json"
	"io"
	//uses "github.com/sea350/ustart_go/backend/uses"
)

type form struct {
	Email      string `json:"Email"`
	Username   string `json:"Username"`
	Password   string `json:"Password"`
	Fname      string `json:"Fname"`
	Lname      string `json:"Lname"`
	University string `json:"University"`
}

func parseRequest(rawData io.ReadCloser) (form, error) {
	var data form
	err := json.NewDecoder(rawData).Decode(&data)
	return data, err
}
