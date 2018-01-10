package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

func ReindexMonologue(eclient *elastic.Client, newMonologue types.Monologue, monoID string) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.MonologueIndex).Do(ctx)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.MonologueIndex).
		Type(globals.MonologueType).
		Id(monoID).
		BodyJson(newMonologue).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
