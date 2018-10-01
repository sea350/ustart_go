package globals

import (
	"context"
	"strings"

	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewClient(elastic.SetURL("http://localhost:9200"))

func DeleteByID(eclient *elastic.Client, docID string, index string) error {

	var idx string
	var typ string
	switch strings.ToLower(index) {
	case "user":
		idx = UserIndex
		typ = UserType
	case "project":
		idx = ProjectIndex
		typ = ProjectType
	case "entry":
		idx = EntryIndex
		typ = EntryType
	case "follow":
		idx = FollowIndex
		typ = FollowType
	case "convo":
		idx = ConvoIndex
		typ = ConvoType
	case "proxymsg":
		idx = ProxyMsgIndex
		typ = ProxyMsgType
	case "msg":
		idx = MsgIndex
		typ = MsgType
	case "widget":
		idx = WidgetIndex
		typ = WidgetType

	}
	ctx := context.Background()
	_, err := eclient.Delete().
		Index(idx).
		Type(typ).
		Id(docID).
		Do(ctx)

	return err
}
