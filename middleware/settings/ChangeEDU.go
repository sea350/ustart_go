package settings

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	uses "github.com/sea350/ustart_go/uses"
)

//ChangeEDU ...
func ChangeEDU(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	typeAcc := r.FormValue("type_select")
	i, err2 := strconv.Atoi(typeAcc)
	if err2 != nil {
		fmt.Println(err2)
	}
	highschoolName := r.FormValue("schoolname")
	highschoolGrad := r.FormValue("highSchoolGradDate")
	uniName := r.FormValue("universityName")
	var major []string
	major = append(major, r.FormValue("majors"))
	//	Year := r.FormValue("year")
	gradDate := r.FormValue("uniGradDate")

	var minor []string

	err := uses.ChangeEducation(eclient, session.Values["DocID"].(string), i, highschoolName, highschoolGrad, uniName, gradDate, major, minor)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	http.Redirect(w, r, "/Settings/#educollapse", http.StatusFound)
	return
}
