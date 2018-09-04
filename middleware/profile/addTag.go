package profile

import (
	"html"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/microcosm-cc/bluemonday"
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

	p := bluemonday.UGCPolicy()

	tags := strings.Split(r.FormValue("skillArray"), `","`)
	var sanitizedTags []string
	for idx := range tags {
		sanitizedTags = append(sanitizedTags, p.Sanitize(tags[idx]))
	}
	// htmlTags := p.Sanitize(tags)
	// fmt.Println("formvalue", r.FormValue("skillArray"))
	// fmt.Println("tags", tags)
	//Dont write floating debug text
	if len(sanitizedTags) > 0 {
		sanitizedTags[0] = strings.Trim(sanitizedTags[0], `["`)
		sanitizedTags[len(sanitizedTags)-1] = strings.Trim(sanitizedTags[len(sanitizedTags)-1], `"]`)
	}

	for i := range tags {
		sanitizedTags[i] = html.EscapeString(sanitizedTags[i])
		// fmt.Println("tag", i, tags[i])
	}

	err := post.UpdateUser(client.Eclient, ID, "Tags", sanitizedTags)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
