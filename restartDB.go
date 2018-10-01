package main

import (
	"context"
	"fmt"
	"log"
	"os"

	globals "github.com/sea350/ustart_go/backend/globals"
)

//var eclient, _ = elastic.NewSimpleClient(elastic.SetURL("http://localhost:9200"))

const usrMapping = `
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

const widgetMapping = `
{
    "mappings":{
        "User":{
            "properties":{
                "UserID":{
                    "type":"keyword"
                },
				"Classification":{
					"type":"keyword"
				}

                
            }
        }
    }
}`

const projMapping = `
{
    "mappings":{
        "Project":{
            "properties":{
                "URLName":{
					"type":"keyword",
					
				},
            }
        }
    }
}`

//RestartDB ... deletes all indexes  in ES and recreates them
func RestartDB() {

	ctx := context.Background()
	//deleteIndex, err := eclient.DeleteIndex(globals.ChatIndex).Do(ctx)
	deleteIndex, err := eclient.DeleteIndex(globals.EntryIndex).Do(ctx)
	deleteIndex, err = eclient.DeleteIndex(globals.ProjectIndex).Do(ctx)
	deleteIndex, err = eclient.DeleteIndex(globals.WidgetIndex).Do(ctx)
	deleteIndex, err = eclient.DeleteIndex(globals.UserIndex).Do(ctx)
	if err != nil {
		// Handle error
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	} else {
		fmt.Println("S U C C")
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
	}

	//ctx = context.Background()
	createIndex, err := eclient.CreateIndex(globals.UserIndex).BodyString(usrMapping).Do(ctx)
	if err != nil {
		// Handle error
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	createIndex, err = eclient.CreateIndex(globals.ProjectIndex).BodyString(projMapping).Do(ctx)
	if err != nil {
		// Handle error
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	createIndex, err = eclient.CreateIndex(globals.WidgetIndex).BodyString(widgetMapping).Do(ctx)
	if err != nil {
		// Handle error
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	createIndex, err = eclient.CreateIndex(globals.EntryIndex).Do(ctx)
	if err != nil {
		// Handle error
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	if !createIndex.Acknowledged {
		// Not acknowledged
	}

}

/*
func main() {
	RestartDB()
}
*/
