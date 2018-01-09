package inbox

import (
	"net/http"
	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
    "github.com/gorilla/sessions"
    client "github.com/sea350/ustart_go/middleware/clientstruct"

)

/* The following 2 lines do not need to be repeated in each subfolder. They can be separated just like everything else.
The first represents the connection to the eS cluster. The second corresponds to a session store. This is not a proper
way to handle the session store. */

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // Needs to be redone

func Inbox(w http.ResponseWriter, r *http.Request){
	// Session called session_please is retreived if it exists
	session, _ := store.Get(r, "session_please")
	// check if docid exists within the session note: there is inconsistency with checking docid/username. 
	test1, _ := session.Values["DocID"]
	// if no docid then redirect to main page 
	if (test1 == nil){
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	// if conditions of redirect are met, the next 4 lines do not execute
	// if docid exists then data is retreived from the uses.UserPage function and rendered accordingly
	userstruct, _, _,_ := uses.UserPage(eclient,session.Values["Username"].(string),session.Values["DocID"].(string))
	cs := client.ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)} 
	client.RenderTemplate(w,"template2-nil",cs)
	client.RenderTemplate(w,"inbox-Nil",cs)
}

