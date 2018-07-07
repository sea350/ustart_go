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

	chatID := r.URL.Path[4:]

	_, err := get.UserByID(client.Eclient, docID.(string))
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
		cs.ErrorOutput = err
		cs.ErrorStatus = true
		client.RenderSidebar(w, r, "template2-nil")
		client.RenderTemplate(w, r, "chat", cs)
		go handleMessages()
		return
	}

	if len(chatID) > 0 {
		if chatID[:1] == `@` {
			//dmUsrID, err := get.IDByUsername(client.Eclient, chatID[1:])
			_, err := get.IDByUsername(client.Eclient, chatID[1:])
			if err != nil {
				log.SetFlags(log.LstdFlags | log.Lshortfile)
				dir, _ := os.Getwd()
				log.Println(dir, err)
				cs.ErrorOutput = err
				cs.ErrorStatus = true
				client.RenderSidebar(w, r, "template2-nil")
				client.RenderTemplate(w, r, "chat", cs)
				go handleMessages()
				return
			}
			//do DM lookup
			//if successfull, pull chat cache
		} else {
			//assume group chat DocID
			//pull chat cache
		}
	}

	//get chat proxy
	//load list of heads
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "chat", cs)
	go handleMessages()
}
