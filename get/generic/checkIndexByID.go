package get

import (
	"context"

	elastic "github.com/olivere/elastic"
)

//CheckIndexByID ...
func CheckIndexByID(eclient *elastic.Client, docID string) (string, error) {
	ctx := context.Background() //intialize context background
	var index string
	searchResult, err := eclient.Get().
		Index("*").
		Id(docID).
		Do(ctx)

	if err != nil {
		return index, err
	}

	index = searchResult.Index

	return index, err

}
