package get

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DMExists ...
//checks to see if a conversation between 2 people already exists
func DMExists(eclient *elastic.Client, eavesdropperOne string, eavesdropperTwo string) (bool, string, error) {

	eavesOne, errOne := get.UserExists(eclient, eavesdropperOne)
	if !eavesOne {
		return false, "", errors.New("E1: Not all participants exist")
	}

	if errOne != nil {
		return false, "", errOne
	}

	eavesTwo, errTwo := get.UserExists(eclient, eavesdropperTwo)
	if !eavesTwo {
		return false, "", errors.New("E2: Not all participants exist")
	}
	if errTwo != nil {
		return false, "", errTwo
	}

	var eaves types.Eavesdropper
	searchMap := make(map[string]types.Eavesdropper)
	searchMap[eavesdropperOne] = eaves
	searchMap[eavesdropperTwo] = eaves
	query := elastic.NewBoolQuery()

	query = query.Should(elastic.NewTermQuery("Eavesdroppers", eavesdropperOne))
	query = query.Should(elastic.NewTermQuery("Eavesdroppers", eavesdropperTwo))

	length := 2
	query = query.Should(elastic.NewTermQuery("EavesCount", length))

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
