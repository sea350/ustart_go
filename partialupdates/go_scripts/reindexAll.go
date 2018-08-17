package partialupdates

import (
	"context"
	"encoding/json"
	"strings"

	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ReindexAll ...
//Reindex all docs in an index
func ReindexAll(eclient *elastic.Client, index string, docType string) error {
	ctx := context.Background()

	maq := elastic.NewMatchAllQuery()
	res, err := eclient.Search().
		Index(index).
		Type(docType).
		Query(maq).
		Do(ctx)

	if err != nil {
		return err
	}

	for _, id := range res.Hits.Hits {
		switch strings.ToLower(docType) {

		case "user":
			data := types.User{}
			err = json.Unmarshal(*id.Source, &data)
			if err != nil {
				return err
			}
		case "project":
			data := types.Project{}
			err = json.Unmarshal(*id.Source, &data)
			if err != nil {
				return err
			}
		case "entry":
			data := types.Entry{}
			err = json.Unmarshal(*id.Source, &data)
			if err != nil {
				return err
			}
		case "proxymsg":
			data := types.ProxyMessages{}
			err = json.Unmarshal(*id.Source, &data)
			if err != nil {
				return err
			}
		case "convo":
			data := types.Conversation{}
			err = json.Unmarshal(*id.Source, &data)
			if err != nil {
				return err
			}
		case "msg":
			data := types.Message{}
			err = json.Unmarshal(*id.Source, &data)
			if err != nil {
				return err
			}

		}

	}

	return nil

}
