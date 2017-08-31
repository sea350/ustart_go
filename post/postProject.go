package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types"github.com/sea350/ustart_go/types"
	"context"
	"errors"
	get"github.com/sea350/ustart_go/get"
)

const PROJ_INDEX = "test-project_data"
const PROJ_TYPE  = "PROJECT"

func IndexProject(eclient *elastic.Client, newProj types.Project)error {
	//ADDS NEW PROJ TO ES RECORDS (requires an elastic client pointer and project type)
	//RETURNS AN error
	ctx := context.Background()

	exists, err := eclient.IndexExists(PROJ_INDEX).Do(ctx)

	if err != nil {return err}

	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(PROJ_INDEX).
		Type(PROJ_TYPE).
		BodyJson(newProj).
		Do(ctx)

	if err != nil {return err}

	return nil
}

func ReindexProject(eclient *elastic.Client, projectID string, projectPage types.Project)error {
	//MODIFIES AN EXISTING PROJ (requires an elastic client pointer, string with project id, 
	//	and the modified project as a project type)
	//RETURNS AN ERROR
	ctx := context.Background()
	
	exists, err := eclient.IndexExists(PROJ_INDEX).Do(ctx)

	if err != nil {return err }
	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(PROJ_INDEX).
		Type(PROJ_TYPE).
		Id(projectID).
		BodyJson(projectPage).
		Do(ctx)

	if err != nil {return err}

	return nil
}

func UpdateProject(eclient *elastic.Client, projectID string, newContent interface{}, field string) error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(PROJ_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err = get.GetProjectById(eclient, projectID)
	if (err!=nil){return err}

	_, err = eclient.Update().
		Index(PROJ_INDEX).
		Type(PROJ_TYPE).
		Id(projectID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	//if err != nil {return err}

	return nil
}

/*
func ModifyName(eclient *elastic.Client, projectID string, newName string)error{

	proj,err := get.GetProjectById(eclient, projectID)
	if (err != nil){return err}

	proj.Name = newName
	
	return UpdateProject(eclient,projectID,proj)
}
*/