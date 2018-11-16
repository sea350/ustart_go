package registration

import (
	"errors"
	"log"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//EmailVerification ... Takes in W and R
//Takes in the form data (email and auth-token) from user and checks if it is correct
//Returns error if false
func EmailVerification(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID != nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := bluemonday.UGCPolicy()
	email := p.Sanitize(r.FormValue("email"))
	emailToken := p.Sanitize(r.FormValue("verifCode"))

	var cs client.ClientSide

	userID, err := get.UserIDByEmail(client.Eclient, email)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		cs.ErrorStatus = true
		cs.ErrorOutput = err
		client.RenderSidebar(w, r, "templateNoUser2")
		client.RenderTemplate(w, r, "reg-got-verified", cs)
		return
	}

	user, err := get.UserByEmail(client.Eclient, email)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		cs.ErrorStatus = true
		cs.ErrorOutput = err
		client.RenderSidebar(w, r, "templateNoUser2")
		client.RenderTemplate(w, r, "reg-got-verified", cs)
		return
	}
	if user.Verified {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	if emailToken == user.AuthenticationCode {
		err = post.UpdateUser(client.Eclient, userID, "Verified", true)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			cs.ErrorStatus = true
			cs.ErrorOutput = err
			client.RenderSidebar(w, r, "templateNoUser2")
			client.RenderTemplate(w, r, "reg-got-verified", cs)
			return
		}
		err = post.UpdateUser(client.Eclient, userID, "AuthenticationCode", nil)
		if err != nil {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			cs.ErrorStatus = true
			cs.ErrorOutput = errors.New("err: middleware/registration/emailVerification line 34")
		}
	} else {
		cs.ErrorStatus = true
		cs.ErrorOutput = errors.New("Verification code is wrong")
	}
	client.RenderSidebar(w, r, "templateNoUser2")
	client.RenderTemplate(w, r, "reg-got-verified", cs)
}
