package settings

import (
	"fmt"
	"net/http"

	"github.com/microcosm-cc/bluemonday"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeContactAndDescription ...
func ChangeContactAndDescription(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()

	var pVIS bool
	var gVIS bool
	var eVIS bool
	phonenumber := p.Sanitize(r.FormValue("pnumber"))

	phonenumbervis := p.Sanitize(r.FormValue("pnumberVis"))
	if phonenumbervis == "True" {
		pVIS = true
	} else {
		pVIS = false
	}
	gender := p.Sanitize(r.FormValue("gender_select"))
	gendervis := p.Sanitize(r.FormValue("gender_selectVis"))
	if gendervis == "True" {
		gVIS = true
	} else {
		gVIS = false
	}
	emailvis := p.Sanitize(r.FormValue("inputEmailVis"))

	if emailvis == "True" {
		eVIS = true
	} else {
		eVIS = false
	}
	description := p.Sanitize(r.FormValue("inputDesc"))
	fmt.Println("Description:", description)
	descriptionrune := []rune(description)
	fmt.Println("Descriptionrune:", descriptionrune)

	userID := session.Values["DocID"].(string)
	err2 := uses.ChangeContactAndDescription(client.Eclient, userID, phonenumber, pVIS, gender, gVIS, eVIS, descriptionrune)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		http.Redirect(w, r, "/Settings/#desccollapse", http.StatusFound)
		return
	}
}
