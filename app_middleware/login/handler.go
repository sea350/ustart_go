package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"github.com/sea350/ustart_go/backend/types"
	uses "github.com/sea350/ustart_go/backend/uses"

	elastic "gopkg.in/olivere/elastic.v5"
)

//var eclient, _ = elastic.NewSimpleClient(elastic.SetURL("http://localhost:9200"))
var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code
//Handler ...
//Login handler
func Handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handling a login request")
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	/*if test1 != nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}*/
	fmt.Println(test1)
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

	//defer json.NewEncoder(w).Encode(resp)

	//Parse request

	data := form{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	var appSessUsr types.AppSessionUser
	resp.updateResp(err, false, appSessUsr)

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

	succ, sessUsr, err := uses.Login(eclient, data.Email, []byte(data.Password), clientIP)

	appSessUsr.FirstName = sessUsr.FirstName
	appSessUsr.LastName = sessUsr.LastName
	appSessUsr.Username = sessUsr.Username
	appSessUsr.Email = sessUsr.Email
	appSessUsr.DocID = sessUsr.DocID

	if !succ {
		fmt.Println("Invalid login")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp.updateResp(errors.New("Password mismatch"), succ, appSessUsr)

		resJson, _ := json.Marshal(resp)
		w.Write(resJson)

	} else {
		fmt.Println("Valid login")
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp.updateResp(err, succ, appSessUsr)
		resJson, errM := json.Marshal(resp)
		if errM != nil {
			fmt.Println(errM)
		}
		w.Write(resJson)

		session.Values["DocID"] = sessUsr.DocID
		fmt.Println("user logged in: " + sessUsr.DocID)
		session.Values["FirstName"] = sessUsr.FirstName
		session.Values["LastName"] = sessUsr.LastName
		session.Values["Email"] = sessUsr.Email
		session.Values["Username"] = sessUsr.Username
		expiration := time.Now().Add((30) * time.Hour)
		cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", Expires: expiration, Path: "/"}

		http.SetCookie(w, &cookie)
		session.Save(r, w)
	}
}
