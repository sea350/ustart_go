package settings 

import (
    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"
)

func ImageUpload(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
    if (test1 == nil){
     	fmt.Println(test1)
     	http.Redirect(w, r, "/~", http.StatusFound)
    }
    r.ParseForm()
    blob := r.FormValue("image-data")
    blob = blob[1:len(blob)]
 //   fmt.Println(blob)

      //fmt.Println(reflect.TypeOf(blob))
    err := uses.ChangeAccountImagesAndStatus(eclient, session.Values["DocID"].(string),blob,true,"hello","Avatar");
    if (err != nil){
    	fmt.Println(err);
    }

    http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)

}

