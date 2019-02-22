package get

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//ProjectFromEvent ...
func ProjectFromEvent(eclient *elastic.Client, eventID string) (types.Project, error) {

	evnt, err := EventByID(eclient, eventID)
	if err != nil {
		return types.Project{}, err
	}

	if evnt.IsProjectHost == false {
		return types.Project{}, errors.New("project not hosting event")
	}

	var proj types.Project
	projectID := evnt.Host
	ctx := context.Background()

	searchResult, err := eclient.Get().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Do(ctx)

	if err != nil {
		return proj, err
	}

	err = json.Unmarshal(*searchResult.Source, &proj)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	return proj, err
}
