package profile

import (
	"fmt"
	"net/http"

	// "github.com/sea350/ustart_go/middleware/stringHTML"

)


func DeleteWallPost(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
	}

	//err := uses.RemoveEntry(eclient,session.Values["DocID"].(string),)

}