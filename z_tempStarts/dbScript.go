package main

import (
	post "github.com/sea350/ustart_go/post/badge"

	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	globals "github.com/sea350/ustart_go/globals"
	postChat "github.com/sea350/ustart_go/post/chat"
	postUser "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

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

const MappingProject = `
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

/*const MappingProject = `
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

func deleteIndex(eclient *elastic.Client, index string) error {

	//fmt.Println(globals.EntryIndex)

	ctx := context.Background()
	log.Println("Current index being deleted:", index)
	deleteIndex, err := eclient.DeleteIndex(index).Do(ctx)
	if err != nil {
		// Handle error
		fmt.Println(err)
		fmt.Println(index)
		return err
	} else {
		fmt.Println(index, "deleted")
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
	}
	return err
}

func startIndex(eclient *elastic.Client, index string) error {
	log.Println("Current index being started:", index)
	mapping := "DNE"
	switch index {
	case globals.UserIndex:
		mapping = globals.MappingUsr
	case globals.ProjectIndex:
		mapping = globals.MappingProject
	case globals.WidgetIndex:
		mapping = globals.MappingWidget
	case globals.EventIndex:
		mapping = eventMapping
	case globals.ConvoIndex:
		mapping = globals.MappingConvo
	case globals.BadgeIndex:
		mapping = globals.MappingBadge
	case globals.FollowIndex:
		mapping = globals.MappingFollow

	}

	if mapping != "DNE" {
		ctx := context.Background()
		_, err := eclient.CreateIndex(index).BodyString(mapping).Do(ctx)

		if err != nil {
			// Handle error

			fmt.Println(err)
			fmt.Println("Could not create", index)
			return err
		} else {
			fmt.Println(index, "created")
		}

	} else {
		ctx := context.Background()
		_, err := eclient.CreateIndex(index).Do(ctx)
		if err != nil {
			// Handle error

			fmt.Println(err)

			fmt.Println("Could not create", index)
			return err

		} else {
			fmt.Println(index, "created")
		}

	}
	return nil

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

var help = make(map[string]string)

func main() {

	//Preload for badge testing
	var ustart types.Badge
	ustart.ID = "USTART"
	ustart.Type = "USTART"
	ustart.ImageLink = "https://s3.amazonaws.com/ustart-default/U_badge.png"
	ustart.Roster = []string{"rr2396@nyu.edu", "sea350@nyu.edu", "yh1112@nyu.edu", "mrb588@nyu.edu", "steven.armanios@nyu.edu"}
	ustart.Tags = []string{"USTARTAdministrator", "USTARTDev"}

	help["help"] = "pretty self-explanatory"
	help["wipe"] = "clears database and restarts all indices"
	help["delete"] = "clears database"
	help["start"] = "starts indices"
	help["deluser"] = "wipes user index"
	help["delproject"] = "wipes project index"
	help["delevent"] = "wipes event index"
	help["delwidget"] = "wipes widget index"
	help["delchat"] = "wipes all chat related indices (Proxy Messages, Messages, and Conversation indices)"
	help["delentries"] = "wipes all entries"
	help["start user"] = "starts user index"
	help["start project"] = "starts project index"
	help["delevent"] = "wipes event index"
	help["delwidget"] = "wipes widget index"
	help["start chat"] = "starts all chat related indices (Proxy Messages, Messages, and Conversation indices)"
	help["stop"] = "end of input"
	//help["remove"] = "removes command from list"
	help["redo"] = "clears current command list"

	indices := []string{}
	commands := []string{}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Welcome to your database service. For help, please input 'help' ")
		text, err := reader.ReadString('\n')
		input := strings.ToLower(text)

		if strings.HasPrefix(input, "stop") {
			fmt.Println("Stopped")
			fmt.Println(commands)
			os.Exit(0)
		} else if err != nil {
			log.Println("Error encountered:", err)

			os.Exit(0)
		} else if strings.HasPrefix(input, "help") {
			for key, val := range help {
				fmt.Println(key, ": ", val)
			}
			commands = append(commands, input)
		} else {
			fmt.Println("Command will be performed")
			commands = append(commands, input)

			if strings.HasPrefix(input, "wipe") {
				indices = append(indices, globals.UserIndex, globals.ProjectIndex, globals.EntryIndex, globals.ConvoIndex, globals.ProxyMsgIndex, globals.MsgIndex, globals.GuestCodeIndex, globals.NotificationIndex, globals.ProxyNotifIndex, globals.WidgetIndex, globals.FollowIndex, globals.ImgIndex, globals.EventIndex, globals.BadgeIndex)
				// delete phase
				for _, index := range indices {

					deleteIndex(eclient, index)
				}

				// restore phase
				for _, index := range indices {

					startIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "delete") {
				indices = append(indices, globals.UserIndex, globals.ProjectIndex, globals.EntryIndex, globals.ConvoIndex, globals.ProxyMsgIndex, globals.MsgIndex, globals.GuestCodeIndex, globals.NotificationIndex, globals.ProxyNotifIndex, globals.WidgetIndex, globals.FollowIndex, globals.ImgIndex, globals.EventIndex, globals.BadgeIndex)
				for _, index := range indices {
					err = deleteIndex(eclient, index)
					if err != nil {
						log.Panicln(err)
						continue
					}
				}
			} else if strings.HasPrefix(input, "start") {
				indices = append(indices, globals.UserIndex, globals.ProjectIndex, globals.EntryIndex, globals.ConvoIndex, globals.ProxyMsgIndex, globals.MsgIndex, globals.GuestCodeIndex, globals.NotificationIndex, globals.ProxyNotifIndex, globals.WidgetIndex, globals.FollowIndex, globals.ImgIndex, globals.EventIndex, globals.BadgeIndex)
				for _, index := range indices {
					err = startIndex(eclient, index)
					if err != nil {
						log.Println(err)
						continue
					}

				}
				_, err := post.IndexBadge(eclient, ustart)
				if err != nil {
					log.Println(err)

				}

			} else if strings.HasPrefix(input, "deluser") {

				indices = append(indices, globals.UserIndex)
				for _, index := range indices {
					deleteIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "delproject") {
				indices = append(indices, globals.ProjectIndex)
				for _, index := range indices {
					deleteIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "delevent") {
				indices = append(indices, globals.EventIndex)
				for _, index := range indices {
					deleteIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "delwidget") {
				indices = append(indices, globals.WidgetIndex)
				for _, index := range indices {
					deleteIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "delentries") {
				indices = append(indices, globals.EntryIndex)
				for _, index := range indices {
					deleteIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "delchat") {
				proxyErr := clearUserProxies(eclient)
				if proxyErr != nil {
					log.Println(proxyErr)
				}
				indices = append(indices, globals.ConvoIndex, globals.ProxyMsgIndex, globals.MsgIndex)
				for _, index := range indices {
					deleteIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "start user") {
				indices = append(indices, globals.UserIndex)
				for _, index := range indices {
					startIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "start project") {
				indices = append(indices, globals.ProjectIndex)
				for _, index := range indices {
					startIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "start event") {
				indices = append(indices, globals.EventIndex)
				for _, index := range indices {
					startIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "start widget") {
				indices = append(indices, globals.WidgetIndex)
				for _, index := range indices {
					startIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "start entries") {
				indices = append(indices, globals.EntryIndex)
				for _, index := range indices {
					startIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "stchat") {

				indices = append(indices, globals.ConvoIndex, globals.ProxyMsgIndex, globals.MsgIndex)
				for _, index := range indices {
					startIndex(eclient, index)
				}
			} else if strings.HasPrefix(input, "redo") {
				commands = []string{}
			} else {
				log.Println("Command invalid")
			}
		}

	}

}
