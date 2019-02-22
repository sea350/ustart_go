package post

import (
	elastic "github.com/olivere/elastic"
	//"errors"

	post "github.com/sea350/ustart_go/post/entry"
	//"golang.org/x/crypto/bcrypt"
)

//UpdateEntry ...
//Sets entries to !Visible for now
func UpdateEntry(eclient *elastic.Client, entryID string, newContent interface{}) error {

	err := post.UpdateEntry(eclient, entryID, "Content", newContent)

	return err

}
