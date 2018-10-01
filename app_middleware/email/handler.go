package email

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	get "github.com/sea350/ustart_go/backend/get/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

//Handler ...
//  Handles registration requests.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING AN EMAIL CHECK REQUEST")
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

	//attempting to catch client IP
	var clientIP string
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Printf("userip: %q is not IP:port\n", r.RemoteAddr)
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Printf("userip: %q is not IP:port\n", r.RemoteAddr)
	} else {
		clientIP = userIP.String()
	}

	fmt.Println(clientIP)
	//func SignUpBasic(eclient *elastic.Client, username string, email string, password []byte, fname string, lname string, country string, state string, city string, zip string, school string, major []string, bday time.Time, currYear string) error {
	inUse, err := get.EmailInUse(eclient, data.Email)

	if err == nil {
		fmt.Println("Valid email")
		resp.updateResp(true, err)
	} else {
		fmt.Println("Taken email")
		resp.updateResp(false, err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp.updateResp(inUse, err)
	resJSON, _ := json.Marshal(resp)
	if false {
		fmt.Println(resJSON)
	}
	// w.Write(resJson)
}
