package app_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("localhost:9200"))

type form struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

//Handler ...
//Login handler
func Handler(w http.ResponseWriter, r *http.Request) {
	resp := setupResp()

	if acrh, ok := r.Header["Access-Control-Request-Headers:"]; ok {
		w.Header().Set("Access-Control-Allow-Origin", acrh[0])

	}

	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := r.Header["Access-Control-Allow-Origin"]; ok {
	} else {
		if _, oko := r.Header["Origin"]; oko {
			w.Header().Set("Access-Control-Allow-Origin", r.Header["Origin"][0])

		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")

		}
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	w.Header().Set("Connection", "Close")

	defer json.NewDecoder(w).Encode(resp)

	//Parse request
	data := form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	fmt.Println("Print request:")
	fmt.Println("%v\n", data)

	resp.updateResp("", 0, err)

	//Get pass from DB
	succ, sessUsr, err := uses.Login(eclient, data.Email, []byte(data.Password))

	if !succ {
		fmt.Println("Invalid login")
		resp.updateResp("", 0, errors.New("Password mismatch"))

	} else {
		fmt.Println("Valid login")
		//resp.updateResp(session.Create(data.Username))
	}
}
