package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"context"

	 
)

const PROJ_INDEX="test-project_data"
const PROJ_TYPE="PROJECT"


func GetProjectFromId(eclient *elastic.Client, projectID string){
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONATAINING
	//		PROJECT DOC ID)
	//RETURNS A types. Project AND AN error
	ctx:=context.Background()
	searchResult, err := eclient.Get().
		Index(PROJ_INDEX).
        Type(PROJ_TYPE).
        Id(projectID).
        Do(ctx)
	
	//if (err!=nil) {return nil, err}

	proj:=searchResult
	return proj, err
	
}