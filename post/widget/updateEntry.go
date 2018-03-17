package post

import (
	elastic "gopkg.in/olivere/elastic.v5"
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
