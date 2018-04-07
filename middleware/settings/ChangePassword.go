package settings 

import (

    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"

)


func ChangePassword(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
		return
    }
    r.ParseForm()
    oldp := r.FormValue("oldpass")
    newp := r.FormValue("confirmpass")
    oldpb := []byte(oldp)
    newpb := []byte(newp)
 //   fmt.Println(blob)

      //fmt.Println(reflect.TypeOf(blob))
    err := uses.ChangePassword(eclient, session.Values["DocID"].(string),oldpb,newpb);
    if (err != nil){
    	fmt.Println(err);
    }
    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
    return

}
