package settings 

import (

    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"

)

func ChangeName(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
		return
    }
    r.ParseForm()
    first := r.FormValue("fname")
    last := r.FormValue("lname")
 //   fmt.Println(blob)

      //fmt.Println(reflect.TypeOf(blob))
    err := uses.ChangeFirstAndLastName(eclient, session.Values["DocID"].(string),first,last);
    if (err != nil){
    	fmt.Println(err);
    }
    http.Redirect(w, r, "/Settings/#namecollapse", http.StatusFound)

}