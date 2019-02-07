package uses

import (
	post "github.com/sea350/ustart_go/post/badge"
	"github.com/sea350/ustart_go/types"

	elastic "gopkg.in/olivere/elastic.v5"
)

//NewBadge...
//Takes care of badge-related modifications and returns relevant tags
func NewBadge(eclient *elastic.Client, badgeID string, badgeType string, badgeTags []string, imageLink string, roster []string) ([]string, error) {
	var newBadge = types.Badge{
		Type:        badgeID,
		ImageLink:   imageLink,
		Roster:      roster,
		Description: desc,
	}

	err := post.IndexBadge(eclient, newBadge)
	if err != nil {
		return nil, err
	}
}
