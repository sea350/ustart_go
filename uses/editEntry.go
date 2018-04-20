package uses

import (
	get "github.com/sea350/ustart_go/get/entry"
	post "github.com/sea350/ustart_go/post/entry"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//EditEntry ...
//Edits an existing widget in the UserWidgets array
func EditEntry(eclient *elastic.Client, entryID string, field string, newVal interface{}) (types.Entry, error) {
	updateErr := post.UpdateEditEntry(eclient, entryID, field, newVal)

	if updateErr != nil {
		panic(updateErr)
	}

	retEnt, err := get.EntryByID(eclient, entryID)
	return retEnt, err

}
