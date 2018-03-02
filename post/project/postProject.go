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
				"Member":{
					"type":"keyword"
				}
            }
        }
    }
}`

//IndexProject ... ADDS NEW PROJECT TO ES RECORDS
//Needs a type Project struct
//returns the new project's id and an error
func IndexProject(eclient *elastic.Client, newProj types.Project) (string, error) {

	tempString := "Project has not been indexed"
	ctx := context.Background()

	exists, err := eclient.IndexExists(globals.ProjectIndex).Do(ctx)

	if err != nil {
		return tempString, err
	}

	if !exists {
		createIndex, Err := eclient.CreateIndex(globals.ProjectIndex).BodyString(projMapping).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(globals.ProjectIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}
	}

	storedProj, err := eclient.Index().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		BodyJson(newProj).
		Do(ctx)

	if err != nil {
		return tempString, err
	}

	return storedProj.Id, nil
}
