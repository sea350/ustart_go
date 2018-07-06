package chat

import (
	"log"
	"net/http"
	"os"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
)

//Page ... draws chat page
func Page(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	cs := client.ClientSide{}

	//chatID := r.URL.Path[4:]

	usr, err := get.UserByID(client.Eclient, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		cs.ErrorOutput = err
		cs.ErrorStatus = true
	} else {
		if usr.ProxyMessagesID == `` {
			//create new proxy
			//index a new proxy
			//update user with new proxy id
		}
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "chat", cs)
	go handleMessages()

}
