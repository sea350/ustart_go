package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"

	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A SETTINGS REQUEST")
	session, _ := store.Get(r, "session_please")
	test1, boo := session.Values["DocID"]

	fmt.Println("LINE 26 bool", boo)

	if test1 == nil {
		fmt.Println("YOU HERE", test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	// Setup the response
	resp := &response{
		Successful: false,

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

	data := form{}

	resp.update(false, json.NewDecoder(r.Body).Decode(&data))
	//	marshalledData, err := json.Marshal(data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	// Validate requestor token
	//valid := true
	//valid, err := session.Validate(data.User, data.Token)

	//resp.update(false, errors.New("Error"))

	fmt.Println("THE INTENT:", data.Intent)
	if test1 == test1 {
		switch data.Intent {
		case "cu":
			fmt.Println("RIGHT INTENT")
			if session.Values["Username"] == data.Username {
				err = uses.ChangeUsername(eclient, session.Values["DocID"].(string), data.Username, data.NewUName)
				resp.update(err == nil, err)

			}
		case "cn":
			if session.Values["Username"] == data.Username {
				err := uses.ChangeFirstAndLastName(eclient, session.Values["DocID"].(string), data.FirstName, data.LastName)
				resp.update(err == nil, err)
			}
		case "cp":
			if session.Values["Username"] == data.Username {
				err = uses.ChangePassword(eclient, session.Values["DocID"].(string), []byte(data.Password), []byte(data.NewPassword))
				resp.update(err == nil, err)
			}
		case "ca":
			if session.Values["Username"] == data.Username {
				err = post.UpdateUser(eclient, session.Values["DocID"].(string), "Avatar", data.Avatar)
				resp.update(err == nil, err)
			}
		case "cb":
			if session.Values["Username"] == data.Username {
				//blob := r.FormValue("banner-data")
				err := post.UpdateUser(eclient, session.Values["DocID"].(string), "Banner", data.Banner)
				resp.update(err == nil, err)
			}
		}
		/*case "gu":
		if session.Values["Username"] == data.Username {
			err = get.
			err = uses.ChangeUsername(eclient, session.Values["DocID"].(string), data.Username, data.NewUName)
			resp.update(err == nil, err, "")
		}*/

	} else {
		resp.update(false, errors.New("Session invalid"))
	}

}
