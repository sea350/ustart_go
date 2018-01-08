package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//DeleteProjectLink ... ADDS A LINK TYPE TO A PROJECT'S QUICKLINKS
//Requires project's docID and a type link
//Returns an error
func DeleteProjectLink(eclient *elastic.Client, projectID string, link types.Link) error {
	ctx := context.Background()
	proj, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	//replace with universal.FindIndex when it works
	index := -1
	for i := range proj.QuickLinks {
		if proj.QuickLinks[i] == link {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("link not found")
	}

	proj.QuickLinks = append(proj.QuickLinks[:index], proj.QuickLinks[index+1:]...)

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"Quicklinks": proj.QuickLinks}).
		Do(ctx)

	return err

}
