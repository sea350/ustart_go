package get

import (
	"context"
	"encoding/json"
	"errors"

	elastic "github.com/olivere/elastic"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
)

//ConvoByID ..
//Retreive a conversation using its ID
func ConvoByID(eclient *elastic.Client, convoID string) (types.Conversation, error) {

	var convo types.Conversation //initialize type chat

	ctx := context.Background()         //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.ConvoIndex).
						Type(globals.ConvoType).
						Id(convoID).
						Do(ctx)

	if err != nil {
		return convo, err
	}

	if !searchResult.Found {
		return convo, errors.New("Conversation not found")
	}

	Err := json.Unmarshal(*searchResult.Source, &convo) //unmarshal type RawMessage into user struct
	if Err != nil {
		return convo, Err
	} //forward error

	return convo, Err

}
