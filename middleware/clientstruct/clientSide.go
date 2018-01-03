package clientstruct 

import(
    types "github.com/sea350/ustart_go/types"
    "net/http"
    "html/template"
)

type ClientSide struct {
	DOCID string 
	FirstName string
    LastName string
    Username string
    ErrorR bool 
    ErrorLogin bool 
    UserInfo types.User
    Class string 
    Birthday string
    ImageCode string
    Description string
    Followers int
    Following int
    Page string
    FollowingStatus string 
    Wall []types.JournalEntry 
}

var templates = template.Must(template.ParseFiles("../../../../../www/ustart.tech/followerlist-nil.html","../../../../../www/ustart.tech/emTee.html","../../../../../www/ustart.tech/wallttt.html","../../../../../www/ustart.tech/wallload-nil.html","../../../../../www/ustart.tech/testimage.html","../../../../../www/ustart.tech/ajax-nil.html","../../../../../www/ustart.tech/Membership-Nil.html","../../../../../www/ustart.tech/settings-Nil.html","../../../../../www/ustart.tech/inbox-Nil.html","../../../../../www/ustart.tech/createProject-Nil.html","../../../../../www/ustart.tech/manageprojects-Nil.html","../../../../../www/ustart.tech/projectsF.html","../../../../../www/ustart.tech/new-reg-nil.html","../../../../../www/ustart.tech/loginerror-nil.html","../../../../../www/ustart.tech/test.html", "../../../../../www/ustart.tech/payment-nil.html","../../../../../www/ustart.tech/templateNoUser2.html","../../../../../www/ustart.tech/profile-nil.html","../../../../../www/ustart.tech/template2-nil.html","../../../../../www/ustart.tech/template-footer-nil.html","../../../../../www/ustart.tech/nil-index2.html","../../../../../www/ustart.tech/regcomplete-nil.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, cs ClientSide) {
    //      fmt.Println("rT called")
    err := templates.ExecuteTemplate(w, tmpl+".html", cs)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }