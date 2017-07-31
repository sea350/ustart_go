package get

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	"context"
	"reflect"
)

const USER_INDEX="test-project_data"
const USER_TYPE="PROJECT"


func GetProjectFromId(eclient *elastic.Client, projectID string)(types.Project, error){
	ctx:=context.Background()
	proj, err := eclient.Get().
		Index(USER_INDEX).
        Type(USER_TYPE).
        Id(projectID).
        Do(ctx)
	}

	return proj, err
	
}