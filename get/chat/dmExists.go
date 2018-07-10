package get

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DMExists ...
//checks to see if a conversation between 2 people already exists
func DMExists(eclient *elastic.Client, eavesdroppers []string) (bool, string, error) {

	if len(eavesdroppers) < 2 {
		return false, "", errors.New("invalid number of chat participants")
	}
	if len(eavesdroppers) > 2 {
		return false, "", nil
	}
	query := elastic.NewBoolQuery()

	for e := range eavesdroppers {
		query = query.Should(elastic.NewTermQuery("Eavesdroppers", eavesdroppers[e]))
	}

	query = query.Should(elastic.NewTermQuery("EavesCount", len(eavesdroppers)))

	ctx := context.Background() //intialize context background
	searchResults, err := eclient.Search().
		Index(globals.ConvoIndex).
		Query(query).
		Pretty(true).
		Do(ctx)

	var chatID string

	if err != nil {
		return false, chatID, err
	}

	exists := searchResults.TotalHits() > 0
	if !exists {
		return exists, chatID, err

	}
	for _, ch := range searchResults.Hits.Hits {
		chatID = ch.Id
	}

	return exists, chatID, errors.New("Conversation exists")

}
