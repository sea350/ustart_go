package settings

import (
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//ChangeName ...
func ChangeName(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	first := html.EscapeString(r.FormValue("fname"))
	last := html.EscapeString(r.FormValue("lname"))
	dob := html.EscapeString(r.FormValue("dob"))

	// if len(dob) == 0 {
	// 	return
	// }
	month, err := strconv.Atoi(dob[5:7])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}
	day, err := strconv.Atoi(dob[8:10])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}
	year, err := strconv.Atoi(dob[0:4])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	bday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	err = uses.ChangeFirstAndLastName(client.Eclient, session.Values["DocID"].(string), first, last, bday)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	http.Redirect(w, r, "/Settings/#namecollapse", http.StatusFound)
	return

}
