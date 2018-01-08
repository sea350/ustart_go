package get

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectByURL ...
// queries ES to get project by URL
func ProjectByURL(eclient *elastic.Client, projectURL string) (types.Project, error) {
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT URL)
	//RETURNS A types.Project AND AN error
	ctx := context.Background()
	termQuery := elastic.NewTermQuery("URL", projectURL)
	searchResult, err := eclient.Search().
		Index(globals.ProjectIndex).
		Query(termQuery).
		Do(ctx)

	var result string
	var proj types.Project
	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		break
	}

	proj, _ = ProjectByID(eclient, result)

	return proj, err

}
