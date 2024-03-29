package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ReindexProject ... REPLACES EXISTING ES DOC
//Specify the docid to be replaced and a type Project struct
//returns an error
func ReindexProject(eclient *elastic.Client, projectID string, projectPage types.Project) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ProjectIndex).Do(ctx)

	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		BodyJson(projectPage).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}
