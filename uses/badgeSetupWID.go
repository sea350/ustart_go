package uses

import (
	"context"
	"log"

	get "github.com/sea350/ustart_go/get/badge"
	getUser "github.com/sea350/ustart_go/get/user"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//BadgeSetupWID ...
//Takes care of badge-related modifications and returns relevant tags
func BadgeSetupWID(eclient *elastic.Client, userID string) ([]string, error) {
	ctx := context.Background()

	usr, err := getUser.UserByID(eclient, userID)

	// var result string
	var tags []string

	for _, id := range usr.BadgeIDs {
		searchResult, err := eclient.Get().
			Index(globals.BadgeIndex).
			Type(globals.BadgeType).
			Id(id).
			Do(ctx)

		if err != nil {
			log.Println(err)

		}

		badge, err := get.BadgeByID(eclient, searchResult.Id)
		if err != nil {
			log.Println(err)
			continue
		} else {
			tags = append(tags, badge.Tags...)

		}

	}

	return tags, err

}
