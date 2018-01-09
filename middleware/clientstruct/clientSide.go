package clientstruct 

import(
    types "github.com/sea350/ustart_go/types"
    "net/http"
    "html/template"
)

/* This struct represents a user state after he/she has logged in. Some fields may no longer be needed 
or are unnecessary.  */ 
type ClientSide struct {
	DOCID string 
	FirstName string
    LastName string
    Username string
    ErrorR bool // registration error likely 
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
/* The following line is how HTML is loaded by our application. Note we need the relative link from the location of GoStart2. */
var templates = template.Must(template.ParseFiles("../../../../../../../www/ustart.tech/followerlist-nil.html","../../../../../../../www/ustart.tech/emTee.html","../../../../../../../www/ustart.tech/wallttt.html","../../../../../../../www/ustart.tech/wallload-nil.html","../../../../../../../www/ustart.tech/testimage.html","../../../../../../../www/ustart.tech/ajax-nil.html","../../../../../../../www/ustart.tech/Membership-Nil.html","../../../../../../../www/ustart.tech/settings-Nil.html","../../../../../../../www/ustart.tech/inbox-Nil.html","../../../../../../../www/ustart.tech/createProject-Nil.html","../../../../../../../www/ustart.tech/manageprojects-Nil.html","../../../../../../../www/ustart.tech/projectsF.html","../../../../../../../www/ustart.tech/new-reg-nil.html","../../../../../../../www/ustart.tech/loginerror-nil.html","../../../../../../../www/ustart.tech/test.html", "../../../../../../../www/ustart.tech/payment-nil.html","../../../../../../../www/ustart.tech/templateNoUser2.html","../../../../../../../www/ustart.tech/profile-nil.html","../../../../../../../www/ustart.tech/template2-nil.html","../../../../../../../www/ustart.tech/template-footer-nil.html","../../../../../../../www/ustart.tech/nil-index2.html","../../../../../../../www/ustart.tech/regcomplete-nil.html"))

/* This function does the actual rendering of HTML pages. Note it takes in a struct (type ClientSide). 
You will need to continually send data to the pages and this is accomplished via the struct. */ 
func RenderTemplate(w http.ResponseWriter, tmpl string, cs ClientSide) {
    err := templates.ExecuteTemplate(w, tmpl+".html", cs)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  }