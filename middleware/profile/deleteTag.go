package profile

import (
	"net/http"
	"strings"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	uses "github.com/sea350/ustart_go/uses"
)

//DeleteTag ...
func DeleteTag(w http.ResponseWriter, r *http.Request) {
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

	deleteTag := r.FormValue("UNKNOWN")

	if ID == r.URL.Path[10:] {
		var newArr []string

		if len(usr.Tags) == 1 {
			err := post.UpdateUser(client.Eclient, ID, "Tags", newArr)
			if err != nil {

				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
			}
			return
		}

		target := -1
		isAllowed, err := uses.TagAllowed(client.Eclient, strings.ToLower(deleteTag))

		if isAllowed {

			for index, tag := range usr.Tags {

				if tag == deleteTag {
					target = index
					break
				}
			}
		}

		if target == -1 {
			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Error: middleware/profile/deleteTag line 54")
			client.Logger.Println("DocID: " + session.Values["DocID"].(string) + " | " + "Deleted object not found")
			newArr = usr.Tags
		} else if (target + 1) < len(usr.Tags) {
			newArr = append(usr.Tags[:target], usr.Tags[(target+1):]...)
		} else {
			newArr = usr.Tags[:target]
		}

		err = post.UpdateUser(client.Eclient, ID, "Tags", newArr)
		if err != nil {

			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: ", err)
		}
	}
}
