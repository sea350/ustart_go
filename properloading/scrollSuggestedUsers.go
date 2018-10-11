package properloading

import (
	"context"
	"io"
	"log"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
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

	followIDs := make([]interface{}, 0)
	for id := range followingUsers {
		followIDs = append([]interface{}{strings.ToLower(id)}, followIDs...)
	}

	followIDs = append([]interface{}{strings.ToLower(userID)}, followIDs...)

	suggestedUserQuery := elastic.NewBoolQuery()
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Tags", tags...))
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Projects.ProjectID", projectIDs...))
	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))

	amt := 1
	if scrollID == `` {
		amt = 3
	}

	searchResults := eclient.Scroll().
		Index(globals.UserIndex).
		Query(suggestedUserQuery).
		Size(amt)

	if len(scrollID) > 0 {
		searchResults = searchResults.ScrollId(scrollID)
	}

	res, err := searchResults.Do(ctx)
	if err != nil && err != io.EOF || res.Hits.TotalHits == 0 {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
		return "", nil, 0, err
	}

	var heads []types.FloatingHead
	for _, hits := range res.Hits.Hits {
		newHead, err := uses.ConvertUserToFloatingHead(eclient, hits.Id)
		if err == nil {
			heads = append(heads, newHead)
		} else {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}

	}

	return res.ScrollId, heads, len(heads), err

}
