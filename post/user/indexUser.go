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
               <!-- "AccCreation":{
                	"type": date"
				},-->
				"FirstName":{
					"type": "keyword"
				},
				"LastName":{
					"type":"keyword"
				}
				"Tags":{
					"type":"keyword"
				}
                
            }
        }
    }
}`

/*TODO: Make this function much better*/

//IndexUser ...
// adds a new user document to the ES cluster
// returns err,string. nil, newID if successful.
func IndexUser(eclient *elastic.Client, newAcc types.User) (error, string) {
	// Check if the index exists
	ctx := context.Background()
	var ID string
	exists, err := eclient.IndexExists(globals.UserIndex).Do(ctx)
	if err != nil {
		return err, ID
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
		return errors.New("Index does not exist"), ID
	}

	// Index the document.
	newUsr, Err := eclient.Index().
		Index(globals.UserIndex).
		Type(globals.UserType).
		BodyJson(newAcc).
		Do(ctx)

	if Err != nil {
		return Err, ID
	}

	return nil, newUsr.Id
}
