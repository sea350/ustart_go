package partialupdates

import (
	"context"
	"errors"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateByQuery ...
//Updates by query
func UpdateByQuery(eclient *elastic.Client, docType string, field string, newContent interface{}) error {
	ctx := context.Background()
	var index string
	var theType string
	switch strings.ToLower(docType) {
	case "user":
		index = globals.UserIndex
		theType = globals.UserType
	case "project":
		index = globals.ProjectIndex
		theType = globals.ProjectType
	case "entry":
		index = globals.EntryIndex
		theType = globals.EntryType
	case "event":
		index = globals.EventIndex
		theType = globals.EventType
	case "convo":
		index = globals.ConvoIndex
		theType = globals.ConvoType
	case "proxymsg":
		index = globals.ProxyMsgIndex
		theType = globals.ProxyMsgType
	default:
		return errors.New("Invalid DOCTYPE")

	}

	exists, err := eclient.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	query := elastic.NewMatchQuery(field, newContent)
	_, err = eclient.UpdateByQuery().
		Refresh("conflicts=proceed").
		Index(index).
		Type(theType).
		Query(query).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
