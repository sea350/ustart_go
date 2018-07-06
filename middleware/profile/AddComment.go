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
	test1, _ := session.Values["Username"]
	if test1 == nil {
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

	_, cmts, err := uses.LoadComments(client.Eclient, postID, 0, -1)
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
	test1, _ := session.Values["Username"]
	fmt.Println("WE ARE IN ADDCOMMENT.GO")
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	r.ParseForm()
	postID := r.FormValue("postID")
	comment := r.FormValue("body")
	contentArray := []rune(comment)
	err := uses.UserReplyEntry(client.Eclient, session.Values["DocID"].(string), postID, contentArray)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	_, cmts, err := uses.LoadComments(client.Eclient, postID, 0, -1)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := json.Marshal(cmts)

	fmt.Fprintln(w, string(data))
}
