package uses

import (
	"time"

	post "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//ChangeFirstAndLastName ...
func ChangeFirstAndLastName(eclient *elastic.Client, userID string, first string, last string, bday time.Time) error {

	err := post.UpdateUser(eclient, userID, "FirstName", first)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "LastName", last)
	if err != nil {
		return err
	}
	err = post.UpdateUser(eclient, userID, "Dob", bday)
	return err
}
