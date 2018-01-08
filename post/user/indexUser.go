package post

import (
	"context"
	"errors"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

const mapping = `
{
    "mappings":{
        "User":{
            "properties":{
                "Email":{
                    "type":"keyword"
                },
                "Username":{
                	"type":"keyword"
                },
                "AccCreation":{
                	"type": date"
				},
				"FirstName":{
					"type": "keyword"
				},
				"LastName":{
					"type":"keyword"
				}

                
            }
        }
    }
}`

/*TODO: Make this function much better*/

//IndexUser ...
// adds a new user document to the ES cluster
// returns err, nil if successful.
func IndexUser(eclient *elastic.Client, newAcc types.User) error {
	// Check if the index exists
	ctx := context.Background()
	exists, err := eclient.IndexExists(globals.UserIndex).Do(ctx)
	if err != nil {
		return err
	}
	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, Err := eclient.CreateIndex(globals.UserIndex).BodyString(mapping).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(globals.UserIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}

		// Return an error saying it doesn't exist
		return errors.New("Index does not exist")
	}

	// Index the document.
	_, Err := eclient.Index().
		Index(globals.UserIndex).
		Type(globals.UserType).
		BodyJson(newAcc).
		Do(ctx)

	if Err != nil {
		return Err
	}

	return nil
}
