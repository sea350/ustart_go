package get

import (
	"context"
	"encoding/json"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ConversationExists ...
//checks to see if a conversation between 2 people already exists
func ConversationExists(eclient *elastic.Client, eavesdroppers []string) (bool, string, error) {

	query := elastic.NewBoolQuery()

	if len(eavesdroppers) < 2 {
		return false, "", errors.New("invalid number of chat participants")
	}

	for e := range eavesdroppers {
		query = query.Should(elastic.NewTermQuery("Eavesdroppers", eavesdroppers[0]))
	}

	ctx := context.Background() //intialize context background
	searchResults, err := eclient.Search().
		Index(globals.UserIndex).
		Query(query).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return false, "", err
	}

	Err := json.Unmarshal(*searchResult.Source, &chat) //unmarshal type RawMessage into user struct
	if Err != nil {
		return false, chat, Err
	} //forward error

	return false, chat, Err

}
