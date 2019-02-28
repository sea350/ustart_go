package profile

import (
	"html"

	"net/http"

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
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	ID := session.Values["DocID"].(string)

	usr, err := get.UserByID(client.Eclient, ID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

	htmlLink := client.SanitizePolicy.Sanitize(r.FormValue("userLink"))
	isValid := uses.ValidLink(htmlLink)
	if len(htmlLink) == 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Link cannot be blank")
		return
	}
	if !isValid {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Invalid link provided")
		return
	}

	cleanTitle := client.SanitizePolicy.Sanitize(r.FormValue("userLinkDesc"))
	if len(cleanTitle) == 0 {
		client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Title cannot be blank")
	}
	usr.QuickLinks = append(usr.QuickLinks, types.Link{Name: html.EscapeString(cleanTitle), URL: html.EscapeString(htmlLink)})

	err = post.UpdateUser(client.Eclient, ID, "QuickLinks", usr.QuickLinks)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
	}

}
