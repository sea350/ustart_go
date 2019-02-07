package uses

import (
	post "github.com/sea350/ustart_go/post/badge"
	"github.com/sea350/ustart_go/types"

	elastic "gopkg.in/olivere/elastic.v5"
)

//NewBadge...
//Takes care of badge-related modifications and returns relevant tags
func NewBadge(eclient *elastic.Client, badgeID string, badgeType string, badgeTags []string, imageLink string, desc string, roster []string) error {
	var newBadge = types.Badge{
		Type:        badgeID,
		ImageLink:   imageLink,
		Roster:      roster,
		Description: desc,
	}

	_, err := post.IndexBadge(eclient, newBadge)
	if err != nil {
		return err
	}
	return err
}
