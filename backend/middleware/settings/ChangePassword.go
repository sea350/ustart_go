package settings

import (
	"log"
	"net/http"
	"os"

	client "github.com/sea350/ustart_go/backend/middleware/client"
	uses "github.com/sea350/ustart_go/backend/uses"
)

//ChangePassword ...
func ChangePassword(w http.ResponseWriter, r *http.Request) {
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
	oldp := r.FormValue("oldpass")
	newp := r.FormValue("confirmpass")
	oldpb := []byte(oldp)
	newpb := []byte(newp)
	err := uses.ChangePassword(client.Eclient, session.Values["DocID"].(string), oldpb, newpb)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	return

}
