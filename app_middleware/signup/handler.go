package signup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	uses "github.com/sea350/ustart_go/uses"
	"golang.org/x/crypto/bcrypt"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//Handler ...
//  Handles registration requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A REGISTER REQUEST")
	// Setup the response
	resp := setupResp()

	if acrh, ok := r.Header["Access-Control-Request-Headers"]; ok {
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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	w.Header().Set("Connection", "Close")

	defer json.NewEncoder(w).Encode(resp)

	// Parse the request, make sure it's A-OK
	data := form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	resp.updateResp(false, err)
	passwordb := []byte(data.Password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(passwordb, bcrypt.DefaultCost)

	//func SignUpBasic(eclient *elastic.Client, username string, email string, password []byte, fname string, lname string, country string, state string, city string, zip string, school string, major []string, bday time.Time, currYear string) error {
	err = uses.SignUpBasic(eclient, data.Username, data.Email, hashedPassword, data.Fname, data.Lname, "", "", "", "", data.University, nil, time.Now(), "", "0")
	if err == nil {
		fmt.Println("Valid signup")
		resp.updateResp(true, err)
	} else {
		fmt.Println("Invalid signup")
		resp.updateResp(false, err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	//resp.updateResp(true, err)
	resJson, _ := json.Marshal(resp)
	if false {
		fmt.Println(resJson)
	}
	// w.Write(resJson)
}
