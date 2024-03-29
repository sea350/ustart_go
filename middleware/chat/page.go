package chat

import (
	
	"net/http"

	"github.com/sea350/ustart_go/uses"

	"github.com/sea350/ustart_go/middleware/client"
)

//Page ... draws chat page
func Page(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	cs := client.ClientSide{}

	chatURL := r.URL.Path[4:]

	valid, _, _, err := uses.ChatVerifyURL(client.Eclient, chatURL, docID.(string))
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}
	if !valid {
		http.Redirect(w, r, "/404/", http.StatusFound)
		return
	}

	client.RenderSidebar(w, r, "template2-nil")
	client.RenderSidebar(w, r, "leftnav-nil")
	client.RenderTemplate(w, r, "chat", cs)
	go handleMessages()
	go HandleChatAlert()
}
