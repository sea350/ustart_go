package uses

import (
	"context"
	"log"
	"strings"

	get "github.com/sea350/ustart_go/get/badge"

	"github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//BadgeSetup ...
//Takes care of badge-related modifications and returns relevant tags
func BadgeSetup(eclient *elastic.Client, email string) ([]string, []string, error) {
	ctx := context.Background()

	//username:= EmailToUsername(email) //for username query
	termQuery := elastic.NewTermQuery("Roster", strings.ToLower(email))
	searchResult, err := eclient.Search().
		Index(globals.BadgeIndex).
		Query(termQuery).
		Do(ctx)

	var badge types.Badge
	if err != nil {
		return nil, nil, err
	}
	var result string
	var tags []string
	var badgeIDs []string
	for _, element := range searchResult.Hits.Hits {

		result = element.Id
		badge, err = get.BadgeByID(eclient, result)
		if err != nil {
			log.Println(err)
			continue
		} else {
			tags = append(tags, badge.Tags...)
			badgeIDs = append(badgeIDs, result)
		}

	}

	return badgeIDs, tags, err

}
