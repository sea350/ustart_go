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

//Handler ...
//Login handler
func Handler(w http.ResponseWriter, r *http.Request) {
	//resp := response{}
	resp := setupResp()

	/*if acrh, ok := r.Header["Access-Control-Request-Headers:"]; ok {
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
	w.Header().Set("Connection", "Close")*/

	defer json.NewEncoder(w).Encode(resp)

	//Parse request

	//byteData := []byte(r.Body)

	data := form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	//byteData, err := ioutil.ReadAll(r.Body)
	//err = json.Unmarshal(byteData, &data)

	fmt.Println("Print request UPDATED:")
	fmt.Println(data.Email)
	fmt.Println(data.Password)

	resp.updateResp("", err)

	fmt.Println("LINE 55, next is err")
	fmt.Println(err)
	fmt.Println("BEFORE LOGIN")
	succ, _, err := uses.Login(eclient, "np1310@nyu.edu", []byte("Ilikedogs1"))

	fmt.Println("AFTER!--Success?", succ)
	//fmt.Println("SESSION USER USERNAME:", sessUsr.Username)

	if !succ {
		fmt.Println("Invalid login")
		resp.updateResp("", errors.New("Password mismatch"))

	} else {
		fmt.Println("Valid login")

		resp.updateResp("np1310@nyu.edu", nil)
	}
}
