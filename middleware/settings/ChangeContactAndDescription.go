package settings 

import (

    "net/http"
    uses "github.com/sea350/ustart_go/uses"
    "fmt"

)


func ChangeContactAndDescription(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if (test1 == nil){
    	fmt.Println(test1)
    	http.Redirect(w, r, "/~", http.StatusFound) }
	r.ParseForm()
	var pVIS bool
	var gVIS bool 
	var eVIS bool 
	phonenumber := r.FormValue("pnumber")
	phonenumbervis := r.FormValue("pnumberVis")
	if (phonenumbervis == "True"){
		pVIS = true 
	}else{
		pVIS = false 
	}
	gender := r.FormValue("gender_select")
	gendervis := r.FormValue("gender_selectVis")
	if (gendervis == "True"){
		gVIS = true 
	}else{
		gVIS = false 
	}
	email := r.FormValue("inputEmail")
	emailvis := r.FormValue("inputEmailVis")
	if (emailvis == "True"){
		eVIS = true 
	}else{
		eVIS = false 
	}
	description := r.FormValue("inputDesc")
	descriptionrune := []rune(description)

	userID := session.Values["DocID"].(string)
    err2 := uses.ChangeContactAndDescription(eclient, userID, phonenumber, pVIS, gender, gVIS, email, eVIS, descriptionrune)   
	if (err2 != nil){
		fmt.Println(err2)
	}   

	if (err2 == nil){
		http.Redirect(w, r, "/profile/"+session.Values["Username"].(string), http.StatusFound)
	}
}