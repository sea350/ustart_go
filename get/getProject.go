package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"context"
	
	"encoding/json"
	
)
const PROJECT_INDEX="test-project_data"
const PROJECT_TYPE="PROJECT"


func GetProjectById(eclient *elastic.Client, projectID string)(types.Project,error){
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT DOC ID)
	//RETURNS A types. Project AND AN error
	var proj types.Project
	
	ctx:=context.Background()

	searchResult, err := eclient.Get().
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		Id(projectID).
		Do(ctx)
	if (err!=nil){return proj,err}
	

	
	err = json.Unmarshal(*searchResult.Source, &proj)
	//if (err!=nil){return proj,err}

	return proj, err
	
}