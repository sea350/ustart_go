package uses

import (
	"log"

	get "github.com/sea350/ustart_go/get/badge"

	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//BadgeSetup ...
//Takes care of badge-related modifications and returns relevant tags
func LoadBadgeProxies(eclient *elastic.Eclient, badgeIDs []string) ([]types.BadgeProxy, err) {
	var bProxies []types.BadgeProxy
	var bProxy types.BadgeProxy
	if len(badgeIDs) == 0 {
		return bProxies, err
	}

	for i := range badgeIDs {
		badge, err := get.BadgeByID(eclient, badgeIDs[i])
		if err != nil {
			log.Println(err)
			continue
		}
		bProxy.Type = badge.Type
		bProxy.Link = badge.ImageLink
		bProxies = append(bProxies, bProxy)

	}

	return bProxies, err

}
