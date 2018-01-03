package inbox

import (
	"net/http"
	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
	"fmt"
    "github.com/gorilla/sessions"
    client "github.com/sea350/ustart_go/middleware/clientstruct"

)

var eclient, err = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
var store = sessions.NewCookieStore([]byte("RIU3389D1")) // code 

func Inbox(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
     if (test1 == nil){
     	fmt.Println(test1)
     http.Redirect(w, r, "/~", http.StatusFound)
       }
    userstruct, _, _,_ := uses.UserPage(eclient,session.Values["Username"].(string),session.Values["DocID"].(string))
    cs := client.ClientSide{UserInfo:userstruct, DOCID:session.Values["DocID"].(string)} 
	client.RenderTemplate(w,"template2-nil",cs)
	client.RenderTemplate(w,"inbox-Nil",cs)
		


}

