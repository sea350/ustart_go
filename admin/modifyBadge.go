package uses

import (
	"log"
	"strings"

	postBadge "github.com/sea350/ustart_go/post/badge"
	post "github.com/sea350/ustart_go/post/user"

	get "github.com/sea350/ustart_go/get/badge"
	getUser "github.com/sea350/ustart_go/get/user"

	elastic "github.com/olivere/elastic"
)

//ModifyBadge...
//Takes care of badge-related modifications and returns relevant tags
func ModifyBadge(eclient *elastic.Client, badgeType string, action string, usrEmail string, newVal string) error {

	usrID, err := getUser.UserIDByEmail(eclient, strings.ToLower(usrEmail))
	if err != nil {
		log.Panicln(err)
		return err
	}
	badge, err := get.BadgeByType(eclient, badgeType)
	if err != nil {
		log.Panicln(err)
		return err
	}
	badgeID := badge.ID
	// var result string

	switch strings.ToLower(action) {
	case "give":
		usr, err := getUser.UserByID(eclient, usrID)
		if err != nil {
			log.Panicln(err)
			return err
		}
		err = post.UpdateUser(eclient, usrID, "BadgeIDs", append(usr.BadgeIDs, badgeID))
		if err != nil {
			log.Panicln(err)
			return err
		}
		err = postBadge.UpdateBadge(eclient, badgeID, "Roster", append(badge.Roster, usrID))
		if err != nil {
			log.Panicln(err)
			return err
		}
	case "image":
		err = postBadge.UpdateBadge(eclient, badgeID, "ImageLink", newVal)
		if err != nil {
			log.Panicln(err)
			return err
		}

	}

	return err

}
