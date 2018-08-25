package settings

import (
	"html"
	"log"
	"net/http"
	"os"

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

	err := uses.ChangeFirstAndLastName(client.Eclient, session.Values["DocID"].(string), first, last)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	http.Redirect(w, r, "/Settings/#namecollapse", http.StatusFound)
	return

}
