package app_middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	uses "github.com/sea350/ustart_go/uses"

	elastic "gopkg.in/olivere/elastic.v5"
)

//var eclient, _ = elastic.NewSimpleClient(elastic.SetURL("http://localhost:9200"))
var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//Handler ...
//Login handler
func Handler(w http.ResponseWriter, r *http.Request) {

	resp := setupResp()

	if acrh, ok := r.Header["Access-Control-Request-Headers:"]; ok {
		w.Header().Set("Access-Control-Allow-Headers", acrh[0])

	}

	w.Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := r.Header["Access-Control-Allow-Origin"]; ok {
		w.Header().Set("Access-Control-Allow-Origin", acao[0])
	} else {
		if _, oko := r.Header["Origin"]; oko {
			w.Header().Set("Access-Control-Allow-Origin", r.Header["Origin"][0])

		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")

		}
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	w.Header().Set("Connection", "Close")

	//defer json.NewEncoder(w).Encode(resp)

	//Parse request

	data := form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	resp.updateResp("", err, false)

	succ, sessUsr, err := uses.Login(eclient, data.Email, []byte(data.Password))

	fmt.Println("SESSUSR", sessUsr)
	fmt.Println(err)

	if !succ {
		fmt.Println("Invalid login")
		resp.updateResp("", errors.New("Password mismatch"), succ)

	} else {
		fmt.Println("Valid login")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp.updateResp(sessUsr.Username, err, succ)
		resJson, _ := json.Marshal(resp)

		w.Write(resJson)
	}
}
