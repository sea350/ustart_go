package get

import (
	"context"

	elastic "gopkg.in/olivere/elastic.v5"
)

//CheckIndexByID ...
func CheckIndexByID(eclient *elastic.Client, docID string) (string, error) {
	ctx := context.Background() //intialize context background
	var index string
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Id(docID).
						Do(ctx)

	if err != nil {
		return index, err
	}

	index = searchResult.Index

	return index, err

}
