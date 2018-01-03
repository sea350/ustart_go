package login
import (

    "net/http"
)

func Logout(w http.ResponseWriter, r *http.Request){
		session, _ := store.Get(r, "session_please")
			test1, _ := session.Values["DocID"]
     if (test1 != nil){
     	session.Options.MaxAge = -1
     	session.Save(r,w)
     	http.Redirect(w, r, "/~", http.StatusFound)

       }
}

