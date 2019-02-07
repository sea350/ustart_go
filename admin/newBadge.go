package uses

import (
	"context"
	"log"

	get "github.com/sea350/ustart_go/get/"
	get "github.com/sea350/ustart_go/get/badge"
	getUser "github.com/sea350/ustart_go/get/user"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//NewBadge...
//Takes care of badge-related modifications and returns relevant tags
func NewBadge(eclient *elastic.Client, badgeID string, badgeType string, badgeTags []string, imageLink string, roster []string) ([]string, error) {
	var newBadge types.Badge{
		Type: badgeID,
		ImageLink :imageLink,
		Roster: roster
		Description: desc string
	}

	err := post.IndexBadge(eclient, newBadge)
	if err != nil {
		return nil, err
	}
}
