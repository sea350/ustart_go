package settings

import (
	"html"

	"net/http"

	"strconv"
	"time"

	"github.com/microcosm-cc/bluemonday"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeName ...
func ChangeName(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	r.ParseForm()

	p := bluemonday.UGCPolicy()

	first := p.Sanitize(r.FormValue("fname"))
	first = html.EscapeString(first)
	// if len(first) < 1 {
	// 			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"First name cannot be blank")

	// }
	last := p.Sanitize(r.FormValue("lname"))
	last = html.EscapeString(last)

	dob := p.Sanitize(r.FormValue("dob"))
	dob = html.EscapeString(dob)
	if len(first) < 1 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "DOB cannot be blank")

	}

	if len(dob) == 0 {
		return
	}
	month, err := strconv.Atoi(dob[5:7])
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		return
	}
	day, err := strconv.Atoi(dob[8:10])
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		return
	}
	year, err := strconv.Atoi(dob[0:4])
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		return
	}

	bday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	err = uses.ChangeFirstAndLastName(client.Eclient, session.Values["DocID"].(string), first, last, bday)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	http.Redirect(w, r, "/Settings/#namecollapse", http.StatusFound)
	return

}
