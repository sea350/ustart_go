package profile

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//DeleteTag ...
func DeleteTag(w http.ResponseWriter, r *http.Request) {
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
		log.Println("Error: middleware/profile/deleteTag line 25")
		log.Println(err)
	}

	deleteTag := r.FormValue("UNKNOWN")

	if ID == r.URL.Path[10:] {
		var newArr []string

		if len(usr.Tags) == 1 {
			err := post.UpdateUser(client.Eclient, ID, "Tags", newArr)
			if err != nil {
				log.Println("Error: middleware/profile/deleteTag line 37")
				log.Println(err)
			}
			return
		}

		target := -1
		for index, tag := range usr.Tags {

			if tag == deleteTag {
				target = index
				break
			}
		}

		if target == -1 {
			log.Println("Error: middleware/profile/deleteTag line 54")
			log.Println("Deleted object not found")
			newArr = usr.Tags
		} else if (target + 1) < len(usr.Tags) {
			newArr = append(usr.Tags[:target], usr.Tags[(target+1):]...)
		} else {
			newArr = usr.Tags[:target]
		}

		err = post.UpdateUser(client.Eclient, ID, "Tags", newArr)
		if err != nil {
			log.Println("Error: middleware/profile/deleteTag line 64")
			log.Println(err)
		}
	}
}
