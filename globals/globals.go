package globals

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

//GuestCodeIndex ...
const GuestCodeIndex = "test-guestcode_data"

//GuestCodeType ...
const GuestCodeType = "GUESTCODE"

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

//ProxyNotifIndex ...
const ProxyNotifIndex = "test-proxynotif_data"

//ProxyNotifType ...
const ProxyNotifType = "PROXYNOTIF"

//NotificationIndex ...
const NotificationIndex = "test-notification_data"

//NotificationType ...
const NotificationType = "NOTIFICATION"

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

const BadgeIndex = "test-badge_index"
const BadgeType = "BADGE"

//MappingBadge...

const MappingBadge = `
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
		"BADGE":{
			"properties":{"
				"Roster":{
					"type":"text",
					"analyzer":"my_analyzer",
					"fields":{
						"raw":{
							"type":"keyword"
							
						}
					}
				},
				"Type":{
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

//MappingUsr ... user mapping
const MappingUsr = `
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

// const MappingUsr = `
// {
//     "mappings":{
//         "USER":{
//             "properties":{
//                 "Email":{
// 					"type":"keyword",

//                 },
//                 "Username":{
// 					"type":"keyword",

//                 },
// 				"FirstName":{
// 					"type": "keyword",

// 				},
// 				"LastName":{
// 					"type":"keyword",

// 				}

//             }
//         }
//     }
// }`

// //MappingWidget ... widget mapping
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

// const MappingProject = `
// {
// 	"settings" :{
// 		"analysis":{
// 			"analyzer" : {
// 				"casesensitive_text":{
// 					"type" : "custom",
// 					"tokenizer": "standard"
// 				}
// 			}
// 		}
// 	},

//     "mappings":{
//         "PROJECT":{
//             "properties":{
// 				"Name":{
// 					"type":"keyword"

// 				},

//                 "URLName":{
// 					"type":"keyword"

// 				},
// 				"Tags":{
// 					"type":"keyword"

// 				}
// 			}

//         }
//     }
// }`

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

//MappingConvo ...
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

const MappingFollow = `
{
   "mappings":{
        "FOLLOW":{
            "properties":{
				"DocID":{
					"type":"keyword"

				}
			}
			
        }
    }
}`
