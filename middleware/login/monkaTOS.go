package login

import (
	"net/http"

	get "github.com/sea350/ustart_go/get/user"
	client "github.com/sea350/ustart_go/middleware/client"
)

//MonkaTOS ... only renders TOS page
func MonkaTOS(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 != nil {
		_, err := get.UserByID(client.Eclient, test1.(string))
		if err != nil {
			cookie := http.Cookie{Name: session.Values["DocID"].(string), Value: "user", MaxAge: -1, Path: "/"}
			http.SetCookie(w, &cookie)
			session.Values = make(map[interface{}]interface{})
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}

	cs := client.ClientSide{}
	client.RenderTemplate(w, r, "templateNoUser2", cs)
	client.RenderTemplate(w, r, "nil-index2", cs)
	client.RenderTemplate(w, r, "tos", cs)
}
