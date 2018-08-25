package profile

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//AddComment ... Iunno
func AddComment(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("followstat")
	postActual := postID[1:]
	comment := r.FormValue("commentz")
	id := r.FormValue("id")
	contentArray := []rune(comment)
	err := uses.UserReplyEntry(client.Eclient, id, postActual, contentArray)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	_, cmts, err := uses.LoadComments(client.Eclient, postID, docID.(string), 0, -1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := json.Marshal(cmts)
	fmt.Println("DATA NEXT:", string(data))
	fmt.Fprintln(w, string(data))
}

//AddComment2 ... Iunno
func AddComment2(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("postID")
	comment := r.FormValue("body")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(comment)
	contentArray := []rune(comment)
	err := uses.UserReplyEntry(client.Eclient, docID.(string), postID, contentArray)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	_, cmts, err := uses.LoadComments(client.Eclient, postID, docID.(string), 0, -1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	data, err := json.Marshal(cmts)

	fmt.Fprintln(w, string(data))
}
