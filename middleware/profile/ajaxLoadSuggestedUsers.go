package profile

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	get "github.com/sea350/ustart_go/get/user"
	"github.com/sea350/ustart_go/globals"
	client "github.com/sea350/ustart_go/middleware/client"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AjaxLoadEntries ... pulls suggested users based on user's projects and shared skills
func AjaxLoadSuggestedUsers(w http.ResponseWriter, r *http.Request) {
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

	//Get tags from user
	tags := make([]interface{}, 0)
	for tag := range myUser.Tags {
		tags = append([]interface{}{strings.ToLower(myUser.Tags[tag])}, tags...)
	}

	//Get mutual project members
	var arrProjects []string
	for _, element := range myUser.Projects {
		arrProjects = append(arrProjects, element.ProjectID)
	}

	memberIDs := make([]interface{}, 0)
	for elements := range arrProjects {
		memberIDs = append([]interface{}{strings.ToLower(arrProjects[elements])}, memberIDs...)
	}

	suggestedUserQuery := elastic.NewBoolQuery()
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Tags", tags...))
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Projects.ProjectID", memberIDs...))

	searchResults, _ := client.Eclient.Search().
		Index(globals.UserIndex).
		Query(suggestedUserQuery).
		Do(ctx)

	data, err := json.Marshal(searchResults)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Println(w, string(data))
}
