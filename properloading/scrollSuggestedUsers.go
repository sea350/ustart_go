package properloading

import (
	"context"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ScrollSuggestedUsers ...
//Scrolls through docs being loaded
func ScrollSuggestedUsers(eclient *elastic.Client, tagArray []string, projects []types.ProjectInfo, followingUsers map[string]bool, userID string, scrollID string) (string, []types.FloatingHead, int, error) {

	ctx := context.Background()
	tags := make([]interface{}, 0)
	for tag := range tagArray {
		tags = append([]interface{}{strings.ToLower(tagArray[tag])}, tags...)
	}

	//Get mutual project members

	projectIDs := make([]interface{}, 0)
	for elements := range projects {
		projectIDs = append([]interface{}{strings.ToLower(projects[elements].ProjectID)}, projectIDs...)
	}

	suggestedUserQuery := elastic.NewBoolQuery()
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Tags", tags...))
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Projects.ProjectID", projectIDs...))

	searchResults, err := client.Eclient.Scroll().
		Index(globals.UserIndex).
		Query(suggestedUserQuery).
		Size(3).
		Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	var heads []types.FloatingHead
	for _, hits := range searchResults.Hits.Hits {
		_, exists := followingUsers[hits.Id]
		if !exists {
			newHead, err := uses.ConvertUserToFloatingHead(eclient, hits.Id)
			if err == nil {
				heads = append(heads, newHead)

			} else {

				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				continue

			}
		}

	}

	return searchResults.ScrollId, heads, int(searchResults.Hits.TotalHits), err

}
