package settings

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/microcosm-cc/bluemonday"

	client "github.com/sea350/ustart_go/middleware/client"
	uses "github.com/sea350/ustart_go/uses"
)

type Major struct {
	List []string
}

//ChangeEDU ...
func ChangeEDU(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	r.ParseForm()
	p := bluemonday.UGCPolicy()
	typeAcc := p.Sanitize(r.FormValue("type_select"))
	i, err2 := strconv.Atoi(typeAcc)
	if err2 != nil {
		fmt.Println(err2)
	}
	highschoolName := p.Sanitize(r.FormValue("schoolname"))
	highschoolGrad := p.Sanitize(r.FormValue("highSchoolGradDate"))
	uniName := p.Sanitize(r.FormValue("universityName"))
	var major []string

	var m Major
	err := json.Unmarshal([]byte(r.FormValue("majors")), &m.List)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	for i := range m.List {
		m.List[i] = p.Sanitize(m.List[i])
		m.List[i] = html.EscapeString(m.List[i])
	}

	major = append(major, r.FormValue("majors"))
	//	Year := r.FormValue("year")
	gradDate := p.Sanitize(r.FormValue("uniGradDate"))

	var minor []string

	err = uses.ChangeEducation(client.Eclient, session.Values["DocID"].(string), i, highschoolName, highschoolGrad, uniName, gradDate, major, minor)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	http.Redirect(w, r, "/Settings/#educollapse", http.StatusFound)
	return
}
