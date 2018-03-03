package post

import (
	"context"

	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

const projMapping = `
{
    "mappings":{
        "Project":{
            "properties":{
                "URLName":{
                    "type":"keyword"
				},
            }
        }
    }
}`

//IndexProject ... ADDS NEW PROJECT TO ES RECORDS
//Needs a type Project struct
//returns the new project's id and an error
func IndexProject(eclient *elastic.Client, newProj types.Project) (string, error) {

	// Check if the index exists
	ctx := context.Background()
	var ID string
	exists, err := eclient.IndexExists(globals.ProjectIndex).Do(ctx)
	if err != nil {
		return ID, err
	}
	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, Err := eclient.CreateIndex(globals.ProjectIndex).BodyString(projMapping).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(globals.ProjectIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}

		// Return an error saying it doesn't exist
		//return ID, errors.New("Index does not exist")
	}

	// Index the document.
	createdProj, Err := eclient.Index().
		Index(globals.ProjectIndex).
		Type("Project").
		BodyJson(newProj).
		Do(ctx)

	if Err != nil {
		return ID, Err
	}

	return createdProj.Id, nil
}
