package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateProject ... UPDATES A SINGLE FEILD IN AN EXISTING ES DOC
//Requires the docID, feild to be modified, and the new content
//Returns an error
func UpdateProject(eclient *elastic.Client, projectID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(projectIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	genericProjectUpdateLock.Lock()
	defer genericProjectUpdateLock.Unlock()

	_, err = get.ProjectByID(eclient, projectID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	//if err != nil {return err}

	return nil
}
