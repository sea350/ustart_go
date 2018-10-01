package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	get "github.com/sea350/ustart_go/backend/get/user"
	"github.com/sea350/ustart_go/backend/globals"
	client "github.com/sea350/ustart_go/backend/middleware/client"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AjaxLoadProjects ... Pulls suggested projects based on skills required
func AjaxLoadSuggestedProjects(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	ID := session.Values["DocID"].(string)
	ctx := context.Background()
	myUser, err := get.UserByUsername(client.Eclient, ID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	arrSkills := make([]interface{}, 0)
	for element := range myUser.Tags {
		arrSkills = append([]interface{}{strings.ToLower(myUser.Tags[element])}, arrSkills...)
	}

	suggestedProjectQuery := elastic.NewBoolQuery()
	suggestedProjectQuery = suggestedProjectQuery.Should(elastic.NewTermsQuery("ListNeeded", arrSkills...))

	searchResults, _ := client.Eclient.Search().
		Index(globals.ProjectIndex).
		Query(suggestedProjectQuery).
		Do(ctx)

	data, err := json.Marshal(searchResults)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Println(w, string(data))
}
