package profile

import (
	"encoding/json"
	"fmt"
	
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//AddComment ... Iunno
func AddComment(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()

	r.ParseForm()
	postID := r.FormValue("followstat")
	postActual := postID[1:]
	comment := p.Sanitize(r.FormValue("commentz"))
	id := r.FormValue("id")
	contentArray := []rune(comment)
	err := uses.UserReplyEntry(client.Eclient, id, postActual, contentArray)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	_, cmts, err := uses.LoadComments(client.Eclient, postID, docID.(string), 0, -1)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	data, err := json.Marshal(cmts)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	fmt.Fprintln(w, string(data))
}

//AddComment2 ... Iunno
func AddComment2(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()

	r.ParseForm()
	postID := r.FormValue("postID")
	comment := p.Sanitize(r.FormValue("body"))

	contentArray := []rune(comment)
	err := uses.UserReplyEntry(client.Eclient, docID.(string), postID, contentArray)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	_, cmts, err := uses.LoadComments(client.Eclient, postID, docID.(string), 0, -1)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	data, err := json.Marshal(cmts)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}
	fmt.Fprintln(w, string(data))
}
