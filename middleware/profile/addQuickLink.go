package profile

import (
	"html"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//AddQuickLink ...
func AddQuickLink(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	ID := session.Values["DocID"].(string)

	usr, err := get.UserByID(client.Eclient, ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	p := bluemonday.UGCPolicy()

	htmlLink := p.Sanitize(r.FormValue("userLink"))
	isValid := uses.ValidLink(htmlLink)
	if len(htmlLink) == 0 {
		log.Println("Link cannot be blank")
		return
	}
	if !isValid {
		log.Println("Invalid link provided")
		return
	}
	if len(r.FormValue("userLinkDesc")) == 0 {
		log.Println("Title cannot be blank")
	}

	cleanTitle := p.Sanitize(r.FormValue("userLinkDesc"))
	usr.QuickLinks = append(usr.QuickLinks, types.Link{Name: html.EscapeString(cleanTitle), URL: html.EscapeString(htmlLink)})

	err = post.UpdateUser(client.Eclient, ID, "QuickLinks", usr.QuickLinks)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

}
