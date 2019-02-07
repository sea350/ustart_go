package uses

import (
	"github.com/sea350/ustart_go/post/user"
	postBadge"github.com/sea350/ustart_go/post/badge"
	"context"
	"log"

	get "github.com/sea350/ustart_go/get/badge"
	getUser "github.com/sea350/ustart_go/get/user"

	"github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
	"os"
)

//ModifyBadge...
//Takes care of badge-related modifications and returns relevant tags
func ModifyBadge(eclient *elastic.Client, badgeType string, action string, usrEmail string, newVal string) ( error) {
	ctx := context.Background()
	usrID, err := getUser.UserIDByEmail(eclient, usrEmail)
	if err != nil {
		log.Panicln(err)
		return err
	}
	badge, err := get.BadgeByType(eclient, badgeType)
	if err != nil {
		log.Panicln(err)
		return err
	}
	badgeID := badge.Id
	// var result string


	switch strings.ToLower(action) {
	case "give":
		usr, err := getUser.UserByID(eclient, usrID)
		if err != nil {
			log.Panicln(err)
			return err
		}
		err = post.UpdateUser(eclient, usrID, "BadgeIDs", append(usr.BadgeIDs, badgeID) 
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

	return err



}
