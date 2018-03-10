package get

import (
	"context"
	"fmt"

	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ProjectByID ...
func ProjectByID(eclient *elastic.Client, projectID string) (types.Project, error) {
	//PULLS FROM ES A PROJECT (REQUIRES AN elastic client pointer AND  A string CONTAINING
	//		PROJECT DOC ID)
	//RETURNS A types. Project AND AN error
	var proj types.Project

	ctx := context.Background()

	searchResult, err := eclient.Get().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Do(ctx)

	if err != nil {
		fmt.Printf("Err: %s\n", err.Error())
		return proj, err
	}

	err = json.Unmarshal(*searchResult.Source, &proj)
	//if (err!=nil){return proj,err}

	fmt.Printf("Err: %s\n", err.Error())

	return proj, err

}
