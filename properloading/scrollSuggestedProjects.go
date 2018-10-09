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

//ScrollSuggestedProjects ...
//Scrolls through docs being loaded
func ScrollSuggestedProjects(eclient *elastic.Client, tagArray []string, projects []types.ProjectInfo, followingProjects map[string]bool, userID string, scrollID string) (string, []types.FloatingHead, int, error) {

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
	for id := range followingProjects {
		followIDs = append([]interface{}{strings.ToLower(id)}, projectIDs...)
	}

	suggestedUserQuery := elastic.NewBoolQuery()
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermsQuery("Tags", tags...))
	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", projectIDs...))
	suggestedUserQuery = suggestedUserQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))
	suggestedUserQuery = suggestedUserQuery.Should(elastic.NewTermQuery("Visisble", true))

	searchResults := eclient.Scroll().
		Index(globals.ProjectIndex).
		Query(suggestedUserQuery).
		Size(1)

	if len(scrollID) > 0 {
		searchResults = searchResults.ScrollId(scrollID)
	}
	res, err := searchResults.Do(ctx)

	if err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
		}
		return "", nil, 0, err
	}

	// var usrIDs []string
	// for _, hits := range res.Hits.Hits {

	// 	_, exists := followingUsers[hits.Id]
	// 	if !exists{
	// 		usrIDs = append(usrIDs, hits.Id)
	// 	}
	// }

	var heads []types.FloatingHead
	for _, hits := range res.Hits.Hits {
		_, exists := followingProjects[hits.Id]
		if !exists && hits.Id != userID {
			newHead, err := uses.ConvertProjectToFloatingHead(eclient, hits.Id)
			if err == nil {
				heads = append(heads, newHead)

			} else {

				log.SetFlags(log.LstdFlags | log.Lshortfile)
				log.Println(err)
				continue

			}
		} else {
			return ScrollSuggestedProjects(eclient, tagArray, projects, followingProjects, userID, res.ScrollId)

		}

	}

	return res.ScrollId, heads, len(heads), err

}
