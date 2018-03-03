package uses

import (
	get "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ConvertToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertToFloatingHead(eclient *elastic.Client, userDocID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	usr, err := get.UserByID(eclient, userDocID)
	if err != nil {
		panic(err)
	}

	head.FirstName = usr.FirstName
	head.LastName = usr.LastName
	head.Image = usr.Avatar
	head.Username = usr.Username

	return head, err
}
