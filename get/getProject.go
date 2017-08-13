package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"context"
	"reflect"
	"encoding/json"
	
)
const PROJECT_INDEX="test-project_data"
const PROJECT_TYPE="PROJECT"


func GetProjectFromId(eclient *elastic.Client, projectID string){
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT DOC ID)
	//RETURNS A types. Project AND AN error
	ctx:=context.Background()
	searchResult, err := eclient.Get().
		Index(PROJECT_INDEX).
        Type(PROJECT_TYPE).
        Id(projectID).
        Do(ctx)
	
    var proj types.Project
	
	Err:= json.Unmarshal(*searchResult.Source, &proj)
	if (Err!=nil){return proj,err}

	return proj, Err
	
}