package email

import (
	"encoding/json"
	"io"
	//uses "github.com/sea350/ustart_go/uses"
)

type form struct {
	Email string `json:"Email"`
}

func parseRequest(rawData io.ReadCloser) (form, error) {
	var data form
	err := json.NewDecoder(rawData).Decode(&data)
	return data, err
}
