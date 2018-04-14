package uses

import (
	post "github.com/sea350/ustart_go/post/user"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChangeUsername ...
func ChangeUsername(eclient *elastic.Client, userID string, newUsername string) error {
	return post.UpdateUser(eclient, userID, "Username", newUsername)

}
