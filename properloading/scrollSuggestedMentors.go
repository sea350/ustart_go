package properloading

import (
	"context"
	"io"
	"log"
	"strings"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//ScrollSuggestedMentors ...
//Scrolls through docs being loaded
func ScrollSuggestedMentors(eclient *elastic.Client, class int, tagArray []string, projects []types.ProjectInfo, followingUsers map[string]bool, userID string, scrollID string) (string, []types.FloatingHead, int, error) {

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

	followingUsers[userID] = true
	followIDs := make([]interface{}, 0)
	for id := range followingUsers {
		followIDs = append([]interface{}{id}, followIDs...)
	}

	suggestedUserQuery := elastic.NewBoolQuery()
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Tags", tags...))
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Projects.ProjectID", projectIDs...))
	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Visible", true))
	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermQuery("_id", userID))
	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Status", true))
	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Verified", true))
	// suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))

	suggestedUserQuery = suggestedUserQuery.Must(elastic.NewTermQuery("Class", 5))

	searchResults := eclient.Scroll().
		Index(globals.UserIndex).
		Query(suggestedUserQuery).
		Size(1).
		Sort("_score", false)

	if len(scrollID) > 0 {
		searchResults = searchResults.ScrollId(scrollID)
	}

	res, err := searchResults.Do(ctx)
	if !(err == io.EOF && res != nil) && err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
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
