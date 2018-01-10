package post

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

func IndexMonologue(eclient *elastic.Client, newMonologue types.Monologue) (string, error) {
	//ADDS NEW MOLOGUE TO ES RECORDS (requires an elastic client and a Monologue type)
	//RETURNS AN error and the new mono's ID IF SUCESSFUL error = nil
	ctx := context.Background()
	var monoID string

	idx, Err := eclient.Index().
		Index(globals.MonologueIndex).
		Type(globals.MonologueType).
		BodyJson(newMonologue).
		Do(ctx)

	if Err != nil {
		return monoID, Err
	}
	monoID = idx.Id

	return monoID, nil
}
