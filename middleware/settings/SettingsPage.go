package settings

import (
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

//Settings ...
func Settings(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	_, err := get.UserByID(client.Eclient, test1.(string))
	if err != nil {
		cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", MaxAge: -1, Path: "/"}
		http.SetCookie(w, &cookie)
		session.Values = make(map[interface{}]interface{})
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	userstruct, _, _, _, _ := uses.UserPage(client.Eclient, session.Values["Username"].(string), session.Values["DocID"].(string))
	cs := client.ClientSide{UserInfo: userstruct, DOCID: session.Values["DocID"].(string)}
	client.RenderSidebar(w, r, "template2-nil")
	client.RenderTemplate(w, r, "settings-Nil", cs)

}
