package post

import (
	"context"
	"errors"

	get "github.com/sea350/ustart_go/get/project"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//AppendProjectLink ... ADDS A LINK TYPE TO A PROJECT'S QUICKLINKS
//Requires project's docID and a type link
//Returns an error
func AppendProjectLink(eclient *elastic.Client, projectID string, link types.Link) error {
	ctx := context.Background()

	proj, err := get.ProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	} //return error if nonexistent

	proj.QuickLinks = append(proj.QuickLinks, link) //retreive project

	_, err = eclient.Update().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Id(projectID).
		Doc(map[string]interface{}{"QuickLinks": proj.QuickLinks}). //update project doc with new link appended to array
		Do(ctx)

	return err

}
