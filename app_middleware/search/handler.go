package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	search "github.com/sea350/ustart_go/backend/search"
	types "github.com/sea350/ustart_go/backend/types"

	elastic "gopkg.in/olivere/elastic.v5"
)

type FloatingSearchHead struct {
	Username string `json:"Username"`
	//for projects Username = project URLName
	FirstName string `json:"FirstName"`
	//for projects Firstname = project Name
	LastName string `json:"LastName"`
	Image    string `json:"Image"`
	Followed bool   `json:"Followed"`
	Bio      string `json:"Bio"`
	//for projects Bio = project Description
	DocID          string `json:"DocID"`
	Classification int    `json:"Classification"`
	Notifications  int    `json:"Notifications"`
}

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code

// Handler responds to http requests about content.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLING A SEARCH REQUEST")
	session, _ := store.Get(r, "session_please")
	test1, boo := session.Values["DocID"]

	fmt.Println("LINE 26 bool", boo)

	// if test1 == nil {
	// 	http.Redirect(w, r, "/~", http.StatusFound)
	// 	return
	// }
	fmt.Println(test1)
	// Setup the response
	resp := &response{
		Successful: false,
		Results:    nil,

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

	//resp.update(false, errors.New(" This is an unknown error"), Proj)
	//	marshalledData, err := json.Marshal(data)

	fmt.Println("Obtained following data: ")
	fmt.Printf("%+v\n", data)

	var results []types.FloatingHead
	var searchRes []FloatingSearchHead
	if test1 == test1 {
		switch data.Intent {
		case "usr":
			fmt.Println("SEARCH USERS")
			results, err = search.PrototypeUserSearch(eclient, strings.ToLower(data.Term), int(0), []bool{true, true, true}, nil, nil, nil)
			for _, res := range results {
				var newRes FloatingSearchHead
				newRes.Bio = string(res.Bio)
				newRes.Classification = res.Classification
				newRes.DocID = res.DocID
				newRes.FirstName = res.FirstName
				newRes.Followed = res.Followed
				newRes.Image = res.Image
				newRes.LastName = res.LastName
				newRes.Notifications = res.Notifications
				newRes.Username = res.Username

				searchRes = append(searchRes, newRes)
			}
			resp.update(err == nil, err, searchRes)
		case "proj":
			fmt.Println("SEARCH PROJECTS")
			results, err = search.PrototypeProjectSearch(eclient, strings.ToLower(data.Term), int(0), []bool{true, true, true, false}, nil, nil, nil)

			for _, res := range results {
				var newRes FloatingSearchHead
				newRes.Bio = string(res.Bio)
				newRes.Classification = res.Classification
				newRes.DocID = res.DocID
				newRes.FirstName = res.FirstName
				newRes.Followed = res.Followed
				newRes.Image = res.Image
				newRes.LastName = res.LastName
				newRes.Notifications = res.Notifications
				newRes.Username = res.Username

				searchRes = append(searchRes, newRes)
			}
			resp.update(err == nil, err, searchRes)

		}
	} else {
		resp.update(false, errors.New("Something went wrong"), searchRes)
	}
}
