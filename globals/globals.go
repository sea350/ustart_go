package get

//ProjectIndex ...
const ProjectIndex = "test-project_data"

//ProjectType ...
const ProjectType = "PROJECT"

//EntryIndex ...
const EntryIndex = "test-entry_data"

//EntryType ...
const EntryType = "ENTRY"

//EventIndex ...
const EventIndex = "test-event_data"

//EventType ...
const EventType = "EVENT"

//UserIndex ...
const UserIndex = "test-user_data"

//UserType ...
const UserType = "USER"

//ChatIndex ...
const ChatIndex = "test-chat_data"

//ChatType ...
const ChatType = "CHAT"

//ConvoIndex ...
const ConvoIndex = "test-convo_data"

//ConvoType ...
const ConvoType = "CONVO"

//MsgIndex ...
const MsgIndex = "test-msg_data"

//MsgType ...
const MsgType = "MSG"

//ProxyMsgIndex ...
const ProxyMsgIndex = "test-proxymsg_data"

//ProxyMsgType ...
const ProxyMsgType = "PROXYMSG"

//WidgetIndex ...
const WidgetIndex = "test-widget_data"

//WidgetType ...
const WidgetType = "WIDGET"

//ImgIndex ...
const ImgIndex = "test-img_data"

//ImgType ...
const ImgType = "IMG"

//FollowIndex ...
const FollowIndex = "test-follow_data"

//FollowType ...
const FollowType = "FOLLOW"

//MappingUsr ... user mapping
const MappingUsr = `
{
    "mappings":{
        "USER":{
            "properties":{
                "Email":{
					"type":"keyword",
					
                },
                "Username":{
					"type":"keyword",
					
                },
				"FirstName":{
					"type": "keyword",
					
				},
				"LastName":{
					"type":"keyword",
					
				}
				
                
            }
        }
    }
}`

//MappingWidget ... widget mapping
const MappingWidget = `
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

//MappingProject ... project mapping
const MappingProject = `
{
	"settings" :{
		"analysis":{
			"analyzer" : {
				"casesensitive_text":{
					"type" : "custom",
					"tokenizer": "standard"
				}
			}
		}
	},

    "mappings":{
        "PROJECT":{
            "properties":{
				"Name":{
					"type":"keyword"

				},

                "URLName":{
					"type":"keyword"
					
				},
				"Tags":{
					"type":"keyword"

				}
			}
			
        }
    }
}`

//MappingEvent ... event mapping
const MappingEvent = `
{
	"settings" :{
		"analysis":{
			"analyzer" : {
				"casesensitive_text":{
					"type" : "custom",
					"tokenizer": "standard"
				}
			}
		}
	},


   "mappings":{
        "EVENT":{
            "properties":{
				"Name":{
					"type":"keyword"

				},
                "URLName":{
					"type":"keyword"
					
				},
				"Tags":{
					"type":"keyword"

				}
			}
			
        }
    }
}`

const MappingConvo = `
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
					
				}
			}
        }
    }
}`
