package profile

import (
	"encoding/json"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/microcosm-cc/bluemonday"
	"github.com/sea350/ustart_go/backend/middleware/client"
	post "github.com/sea350/ustart_go/backend/post/user"
)

type TagStruct struct {
	Tags []string
}

//AddTag ...
func AddTag(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	var ts TagStruct
	err := json.Unmarshal([]byte(r.FormValue("skillArray")), &ts.Tags)

	if err != nil {
		log.Println("Could not unmarshal")
		return
	}
	ID := session.Values["DocID"].(string)

	p := bluemonday.UGCPolicy()

	for t := range ts.Tags {
		ts.Tags[t] = p.Sanitize(ts.Tags[t])
		ts.Tags[t] = html.EscapeString(ts.Tags[t])
	}
	// tags := strings.Split(r.FormValue("skillArray"), `","`)
	// var sanitizedTags []string
	// for idx := range tags {
	// 	sanitizedTags = append(sanitizedTags, p.Sanitize(tags[idx]))
	// }
	// htmlTags := p.Sanitize(tags)
	// fmt.Println("formvalue", r.FormValue("skillArray"))
	// fmt.Println("tags", tags)
	//Dont write floating debug text
	// if len(sanitizedTags) > 0 {
	// 	sanitizedTags[0] = strings.Trim(sanitizedTags[0], `["`)
	// 	sanitizedTags[len(sanitizedTags)-1] = strings.Trim(sanitizedTags[len(sanitizedTags)-1], `"]`)
	// }

	// for i := range tags {
	// 	sanitizedTags[i] = html.EscapeString(sanitizedTags[i])
	// 	// fmt.Println("tag", i, tags[i])
	// }

	err = post.UpdateUser(client.Eclient, ID, "Tags", ts.Tags)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
}
