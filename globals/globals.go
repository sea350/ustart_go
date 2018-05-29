package get

//ProjectIndex ...
const ProjectIndex = "test-project_data"

//ProjectType ...
const ProjectType = "PROJECT"

//EntryIndex ...
const EntryIndex = "test-entry_data"

//EntryType ...
const EntryType = "ENTRY"

//UserIndex ...
const UserIndex = "test-user_data"

//UserType ...
const UserType = "USER"

//ChatIndex ...
const ChatIndex = "test-chat_data"

//ChatType ...
const ChatType = "CHAT"

//WidgetIndex ...
const WidgetIndex = "test-widget_data"

//WidgetType ...
const WidgetType = "WIDGET"

//MappingUsr ... user mapping
const MappingUsr = `
{
    "mappings":{
        "USER":{
            "properties":{
                "Email":{
					"type":"keyword",
					"analyzer":"english"
                },
                "Username":{
					"type":"keyword",
					"analyzer":"english"
                },
               <!-- "AccCreation":{
                	"type": date"
				},-->
				"FirstName":{
					"type": "keyword",
					"analyzer":"english"
				},
				"LastName":{
					"type":"keyword",
					"analyzer":"english"
				}
				<!--"Tags":{
					"type":"keyword"-->
				}
                
            }
        }
    }
}`

//MappingWidget ... widget mapping
const MappingWidget = `
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
