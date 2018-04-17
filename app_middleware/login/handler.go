package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	uses "github.com/sea350/ustart_go/uses"

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

	resp.updateResp(err, false, session)

	succ, sessUsr, err := uses.Login(eclient, data.Email, []byte(data.Password))

	fmt.Println("SESSUSR", sessUsr)
	fmt.Println(err)

	if !succ {
		fmt.Println("Invalid login")
		resp.updateResp(errors.New("Password mismatch"), succ, session)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		resJson, _ := json.Marshal(resp)
		w.Write(resJson)

	} else {
		fmt.Println("Valid login")

		resJson, _ := json.Marshal(resp)
		session.Values["DocID"] = sessUsr.DocID
		fmt.Println("user logged in: " + sessUsr.DocID)
		session.Values["FirstName"] = sessUsr.FirstName
		session.Values["LastName"] = sessUsr.LastName
		session.Values["Email"] = sessUsr.Email
		session.Values["Username"] = sessUsr.Username
		expiration := time.Now().Add((30) * time.Hour)
		cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", Expires: expiration, Path: "/"}

		fmt.Println("THE RESPONSE:", resJson)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp.updateResp(err, succ, session)
		w.Write(resJson)
		http.SetCookie(w, &cookie)
		session.Save(r, w)
	}
}
