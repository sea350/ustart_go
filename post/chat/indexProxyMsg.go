package post

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//IndexProxyMsg ...
//Indexes a new proxy message
func IndexProxyMsg(eclient *elastic.Client, newProxyMsg types.ProxyMessages) (string, error) {
	//ADDS NEW CHAT TO ES RECORDS (requires an elastic client and a Chat type)
	//RETURNS AN error and the new chat's ID IF SUCESSFUL error = nil
	ctx := context.Background()
	var proxyMsgID string

	idx, Err := eclient.Index().
		Index(globals.ProxyMsgIndex).
		Type(globals.ProxyMsgType).
		BodyJson(newProxyMsg).
		Do(ctx)

	if Err != nil {
		return msgID, Err
	}
	proxyMsgID = idx.Id

	return proxyMsgID, nil
}
