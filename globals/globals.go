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
					"index" : "not_analyzed"
                },
                "Username":{
					"type":"keyword",
					"index" : "not_analyzed"
                },
               <!-- "AccCreation":{
                	"type": date"
				},-->
				"FirstName":{
					"type": "keyword",
					"index" : "not_analyzed"
				},
				"LastName":{
					"type":"keyword",
					"index" : "not_analyzed"
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
        "Project":{
            "properties":{
                "URLName":{
					"type":"keyword",
					
					
					"analyzer": "casesensitive_text"
				},
				"Tags":{
					"type":"keyword"
				}
			}
			
        }
    }
}`
