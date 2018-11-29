package profile

import (
	"log"
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
)

//DeleteQuickLink ...
//designed for ajax
func DeleteQuickLink(w http.ResponseWriter, r *http.Request) {
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
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	deleteTitle := r.FormValue("userLinkDesc")
	deleteURL := r.FormValue("userLink")

	if deleteURL == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Crucial data was not passed in, now exiting")
		return
	}
	if deleteTitle == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: link title is blank")
	}

	var newArr []types.Link

	if len(usr.QuickLinks) <= 1 {
		err := post.UpdateUser(client.Eclient, ID, "QuickLinks", newArr)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		return
	}

	target := -1
	for index, link := range usr.QuickLinks {

		if link.Name == deleteTitle && link.URL == deleteURL {
			target = index
			break
		}
	}

	if target == -1 {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("Deleted object not found")
		return
	} else if (target + 1) < len(usr.QuickLinks) {
		newArr = append(usr.QuickLinks[:target], usr.QuickLinks[(target+1):]...)
	} else {
		newArr = usr.QuickLinks[:target]
	}

	err = post.UpdateUser(client.Eclient, ID, "QuickLinks", newArr)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	return
}
