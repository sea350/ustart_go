package get

import (
	"context"
	"encoding/json"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProxyMsgByID ..
//Retreive a proxy message using its ID
func ProxyMsgByID(eclient *elastic.Client, proxyMsgID string) (types.ProxyMessages, error) {

	var proxyMsg types.ProxyMessages //initialize type chat

	ctx := context.Background()         //intialize context background
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.ProxyMsgIndex).
						Type(globals.ProxyMsgType).
						Id(proxyMsgID).
						Do(ctx)

	if err != nil {
		return proxyMsg, err
	}
	if !searchResult.Found {
		return proxyMsg, errors.New("Proxy message not found")
	}
	Err := json.Unmarshal(*searchResult.Source, &proxyMsg) //unmarshal type RawMessage into user struct

	return proxyMsg, Err

}
