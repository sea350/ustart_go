package settings

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//UserChangeBirthday ...  changes the user's birthday designed for ajax
func UserChangeBirthday(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		fmt.Println(test1)
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	dob := r.FormValue("dob")

	month, err := strconv.Atoi(dob[5:7])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}
	day, err := strconv.Atoi(dob[8:10])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}
	year, err := strconv.Atoi(dob[0:4])
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return
	}

	bday := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	err = post.UpdateUser(client.Eclient, test1.(string), "Dob", bday)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
}
