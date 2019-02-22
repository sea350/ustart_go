package uses

import (
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	post "github.com/sea350/ustart_go/post/user"
	elastic "github.com/olivere/elastic"
)

//ChangeUsername ...
func ChangeUsername(eclient *elastic.Client, userID string, oldUsername string, newUsername string) error {
	inUse, err := get.UsernameInUse(eclient, newUsername)

	if err != nil {
		return err
	}

	if !inUse {
		err = post.UpdateUser(eclient, userID, "Username", newUsername)
		return err
	}

	return errors.New("Username taken")

}
