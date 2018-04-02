package uses

import (
	post "github.com/sea350/ustart_go/post/entry"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EditWidget ...
//Edits an existing widget in the UserWidgets array
func EditEntry(eclient *elastic.Client, entryID string, field string, newVal interface{}) error {
	updateErr := post.UpdateEntry(elastic, entryID)

	if updateErr != nil {
		panic(updateErr)
	}

	return updateErr

}