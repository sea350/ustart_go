package main

import (
	"context"
	"fmt"

	globals "github.com/sea350/ustart_go/globals"
	postChat "github.com/sea350/ustart_go/post/chat"
	postUser "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL("https://vpc-ustart-es-ho7jd4ahrgusb6zp2j6qvecvtu.us-east-1.es.amazonaws.com"))

const usrMapping = `
{
	"settings": {
		"analysis": {
		   "analyzer": {
			  "my_analyzer": {
				 "type": "custom",
				 "filter": [
					"lowercase"
				 ],
				 "tokenizer": "whitespace"
			  }
		   }
		}
	 },

    "mappings":{
        "USER":{
            "properties":{
                "Email":{
                    "type":"keyword"
				},
                "Username":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}
                },
				"FirstName":{
					"type": "text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}
				},
				"LastName":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}
				},
				"Tags":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}
				}
				 
                
            }
        }
    }
}`

// "casesensitive_text":{
// 	"type" : "custom",
// 	"tokenizer": "standard"
// },

const projMapping = `
{
	"settings": {
		"analysis": {
		   "analyzer": {
			  "my_analyzer": {
				 "type": "custom",
				 "filter": [
					"lowercase"
				 ],
				 "tokenizer": "whitespace"
			  }
		   }
		}
	 },


    "mappings":{
        "PROJECT":{
            "properties":{
				"Name":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}					 
				},
                "URLName":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}		
					
				},
				"Tags":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}
				}
			}
        }
    }
}`

const eventMapping = `
{
	"settings": {
		"analysis": {
		   "analyzer": {
			  "my_analyzer": {
				 "type": "custom",
				 "filter": [
					"lowercase"
				 ],
				 "tokenizer": "whitespace"
			  }
		   }
		}
	 },


    "mappings":{
        "EVENT":{
            "properties":{
				"Name":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}					 
				},
                "URLName":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}		
					
				},
				"Tags":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}
				}
			}
        }
    }
}`

/*const usrMapping = `
{
    "mappings":{
        "USER":{
            "properties":{
                "Email":{
                    "type":"keyword"
                },
                "Username":{
                	"type":"keyword"
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
}`*/

const widgetMapping = `
{
	
    "mappings":{
        "WIDGET":{
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

/*const projMapping = `
{
    "mappings":{
        "PROJECT":{
            "properties":{
				"Name":{
					"type":"keyword"

				},
                "URLName":{
					"type":"keyword"

				}

            }
        }
    }
}`*/

const convoMapping = `
{
	"settings": {
		"analysis": {
		   "analyzer": {
			  "my_analyzer": {
				 "type": "custom",
				 "filter": [
					"lowercase"
				 ],
				 "tokenizer": "whitespace"
			  }
		   }
		}
	 },


    "mappings":{
        "CONVO":{
            "properties":{
				"Title":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}					 
				},
                "RefProject":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}		
					
				},
				"Eavesdroppers":{
					"DocID":{
						"type":"keyword",
						"analyzer":"not_analyzed"
						
					}
				}
			}
        }
    }
}`

func deleteIndex(eclient *elastic.Client, index string) {

	fmt.Println(globals.EntryIndex)

	ctx := context.Background()
	deleteIndex, err := eclient.DeleteIndex(index).Do(ctx)
	if err != nil {
		// Handle error
		fmt.Println(err)
		fmt.Println(index)
	} else {
		fmt.Println(index, "deleted")
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
	}
}

func startIndex(eclient *elastic.Client, index string) {
	fmt.Println(index)
	mapping := "DNE"
	switch index {
	case globals.UserIndex:
		mapping = usrMapping
	case globals.ProjectIndex:
		mapping = projMapping
	case globals.WidgetIndex:
		mapping = globals.MappingWidget
	case globals.EventIndex:
		mapping = eventMapping
	case globals.ConvoIndex:
		mapping = globals.MappingConvo

	case globals.FollowIndex:
		mapping = globals.MappingFollow

	}

	if mapping != "DNE" {
		ctx := context.Background()
		_, err := eclient.CreateIndex(index).BodyString(mapping).Do(ctx)
		fmt.Println("Mapping:")
		fmt.Println(mapping)
		if err != nil {
			// Handle error
			fmt.Println("Mapping:")
			fmt.Println(mapping)
			fmt.Println(err)
			fmt.Println("Could not create", index)
		} else {
			fmt.Println(index, "created")
		}

	} else {
		ctx := context.Background()
		_, err := eclient.CreateIndex(index).Do(ctx)
		if err != nil {
			// Handle error
			fmt.Println("Mapping:")
			fmt.Println(mapping)
			fmt.Println(err)

			fmt.Println("Could not create", index)

		} else {
			fmt.Println(index, "created")
		}

	}

}

func clearUserProxies(eclient *elastic.Client) error {

	ctx := context.Background()
	maQ := elastic.NewMatchAllQuery()

	res, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maQ).
		Do(ctx)

	if err != nil {
		return err

	}

	for _, usr := range res.Hits.Hits {
		var newProxy types.ProxyMessages
		newProxy.Class = 1
		newProxy.DocID = usr.Id
		proxyID, err := postChat.IndexProxyMsg(eclient, newProxy)
		if err != nil {
			return err

		}
		err = postUser.UpdateUser(eclient, usr.Id, "ProxyMessagesID", proxyID)
		if err != nil {
			return err

		}
	}
	return nil
}

func main() {
	//indices := []string{globals.UserIndex, globals.ProjectIndex, globals.EntryIndex, globals.WidgetIndex}
	// indices := []string{globals.ConvoIndex, globals.ProxyMsgIndex, globals.MsgIndex} //, globals.ProjectIndex}
	// indices := []string{globals.ConvoIndex}
	// indices := []string{globals.ProxyMsgIndex}
	var indices []string
	indices = append(indices, globals.UserIndex)
	indices = append(indices, globals.ProjectIndex)
	indices = append(indices, globals.EntryIndex)
	indices = append(indices, globals.WidgetIndex)
	indices = append(indices, globals.FollowIndex)
	// // fmt.Println(globals.FollowIndex)
	// indices = append(indices, globals.EventIndex)
	// indices = append(indices, )
	// indices = append(indices, )
	// {globals.MsgIndex, globals.ConvoIndex, globals.ProxyMsgIndex}
	//no chat atm
	indices = append(indices, globals.MsgIndex, globals.ConvoIndex, globals.ProxyMsgIndex)

	indices = append(indices, globals.NotificationIndex, globals.ProxyNotifIndex, globals.GuestCodeIndex)

	// indices = append(indices, globals.GuestCodeIndex)
	// delete phase
	for _, index := range indices {
		deleteIndex(eclient, index)
	}
	// deleteIndex(eclient, globals.ConvoIndex)
	// restore phase
	for _, index := range indices {
		startIndex(eclient, index)
	}

	// err := clearUserProxies(eclient)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
