package uses

import (
	postEntry "github.com/sea350/ustart_go/backend/post/entry"
	elastic "gopkg.in/olivere/elastic.v5"
)

//HideEntry ... Sets an entry to invisible
//Requires the entry's docID
func HideEntry(eclient *elastic.Client, entryID string) error {

	err := postEntry.UpdateEntry(eclient, entryID, "Visible", false)
	//This function is a redirect now
	//it is maintained just in case more things need to be executed when a post is deleted in the future

	return err
}
