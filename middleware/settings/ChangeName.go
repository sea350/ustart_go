package settings

import (
	"log"
	"net/http"
	"os"

	uses "github.com/sea350/ustart_go/uses"
)

//ChangeName ...
func ChangeName(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	first := r.FormValue("fname")
	last := r.FormValue("lname")

	err := uses.ChangeFirstAndLastName(eclient, session.Values["DocID"].(string), first, last)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	http.Redirect(w, r, "/Settings/#namecollapse", http.StatusFound)
	return

}
