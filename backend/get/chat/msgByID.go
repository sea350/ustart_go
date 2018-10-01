package get

import (
	"context"
	"encoding/json"
	"errors"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//MsgByID ..
//Retreive a message using its ID
func MsgByID(eclient *elastic.Client, msgID string) (types.Message, error) {

	var msg types.Message //initialize type chat

	ctx := context.Background()         //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.MsgIndex).
						Type(globals.MsgType).
						Id(msgID).
						Do(ctx)

	if err != nil {
		return msg, err
	}

	if !searchResult.Found {
		return msg, errors.New("Message not found")
	}

	Err := json.Unmarshal(*searchResult.Source, &msg) //unmarshal type RawMessage into user struct
	if Err != nil {
		return msg, Err
	} //forward error

	return msg, Err

}
