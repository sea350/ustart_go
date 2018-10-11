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
		followIDs = append([]interface{}{strings.ToLower(id)}, followIDs...)
	}

	suggQuery := elastic.NewBoolQuery()
	suggQuery = suggQuery.Should(elastic.NewTermsQuery("Tags", tags...))
	suggQuery = suggQuery.Should(elastic.NewTermsQuery("ListNeeded", tags...))
	suggQuery = suggQuery.MustNot(elastic.NewTermsQuery("_id", projectIDs...))
	suggQuery = suggQuery.MustNot(elastic.NewTermsQuery("_id", followIDs...))
	suggQuery = suggQuery.Must(elastic.NewTermQuery("Visisble", true))

	amt := 1
	if scrollID == `` {
		amt = 3
	}

	searchResults := eclient.Scroll().
		Index(globals.ProjectIndex).
		Query(suggQuery).
		Size(amt)

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
		newHead, err := uses.ConvertProjectToFloatingHead(eclient, hits.Id)
		if err == nil {
			heads = append(heads, newHead)
		} else if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err)
			continue
		}
	}

	return res.ScrollId, heads, len(heads), err

}
