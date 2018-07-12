package profile

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/user"
)

//AddTag ...
func AddTag(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	ID := session.Values["DocID"].(string)

	testingtag := html.EscapeString("<script>")
	fmt.Println(testingtag)
	tags := strings.Split(r.FormValue("skillArray"), `","`)
	fmt.Println(r.FormValue("skillArray"))
	fmt.Println(tags)

	if len(tags) > 0 {
		tags[0] = strings.Trim(tags[0], `["`)
		tags[len(tags)-1] = strings.Trim(tags[len(tags)-1], `"]`)
	}

	for i := 0; i < len(tags); i++ {
		tags[i] = html.EscapeString(tags[i])
		fmt.Println(tags[i])
	}

	err := post.UpdateUser(client.Eclient, ID, "Tags", tags)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
