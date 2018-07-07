package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/event"
	uses "github.com/sea350/ustart_go/uses"
)

//EventChangeNameAndDescription ...
//For Events
func EventChangeNameAndDescription(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	evntName := r.FormValue("ename")
	evntDesc := []rune(r.FormValue("inputDesc"))
	//   fmt.Println(blob)
	//fmt.Println(projName, projName)

	//fmt.Println(reflect.TypeOf(blob))
	evnt, err := get.EventByID(eclient, r.FormValue("eventID"))
	//TODO: DocID
	err = uses.ChangeEventNameAndDescription(eclient, r.FormValue("eventID"), evntName, evntDesc)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	//TODO: Add in right URL
	http.Redirect(w, r, "/Events/"+evnt.URLName, http.StatusFound)
	return

}
