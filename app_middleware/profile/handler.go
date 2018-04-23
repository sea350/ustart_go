package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A PROFILE VIEW REQUEST")
	session, _ := store.Get(r, "session_please")
	test1, boo := session.Values["DocID"]

	fmt.Println("LINE 26 bool", boo)

	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	Usr := types.User{}
	// Setup the response
	resp := &response{
		Successful: false,
		User:       Usr,

		ErrMsg: errors.New("Unknown failure"),
	}

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

	// Parse the request.
	fmt.Println("Parsing request")

	//var sessUsr types.AppSessionUser
	data := form{}

	err := json.NewDecoder(r.Body).Decode(&data)

	fmt.Println("THIS IS THE DATA:", data)
	if err != nil {
		fmt.Println("error line 70 profile/handler.go")
	}

	Usr, err = get.UserByUsername(eclient, data.Username)
	if err != nil {
		fmt.Println("error line 75 profile/handler.go")
	}

	usrID, err := get.IDByUsername(eclient, data.Username)

	if err != nil {
		fmt.Println("error line 84 profile/handler.go")
	}

	//resp.update(false, errors.New(" This is an unknown error"), Usr)
	//	marshalledData, err := json.Marshal(data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	//valid := true
	//valid, err := session.Validate(data.User, data.Token)

	//resp.update(false, errors.New("Error"))

	if test1 == test1 {
		if data.Intent == "foll" {
			if session.Values["Username"] != data.Username {
				isFollowed, err := uses.IsFollowed(eclient, usrID, session.Values["DocID"].(string))

				if !isFollowed {
					fmt.Println("RIGHT INTENT")
					err = uses.UserFollow(eclient, strings.ToLower(usrID), strings.ToLower(session.Values["DocID"].(string)))
					resp.update(err == nil, err, Usr)
				} else {
					err = uses.UserUnfollow(eclient, usrID, session.Values["DocID"].(string))
					resp.update(err == nil, err, Usr)
				}
			}
		} else {
			resp.update(false, errors.New("Token invalid"), Usr)
		}
	}

}
