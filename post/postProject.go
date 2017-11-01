package post

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types"github.com/sea350/ustart_go/types"
	"context"
	"errors"
	get"github.com/sea350/ustart_go/get"
)

const PROJECT_INDEX = "test-project_data"
const PROJECT_TYPE  = "PROJECT"




const proj_mapping = `
{
    "mappings":{
        "Project":{
            "properties":{
                "URLName":{
                    "type":"keyword"
                }
                
            }
        }
    }
}`

func IndexProject(eclient *elastic.Client, newProj types.Project)(string, error) {
	//ADDS NEW PROJ TO ES RECORDS (requires an elastic client pointer and project type)
	//RETURNS AN error
	tempString := "Project has not been indexed"
	ctx := context.Background()

	exists, err := eclient.IndexExists(PROJECT_INDEX).Do(ctx)

	if err != nil {return tempString, err}

	if !exists {return tempString, errors.New("Index does not exist")}

	storedProj, err := eclient.Index().
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		BodyJson(newProj).
		Do(ctx)

	if err != nil {return tempString, err}

	return storedProj.Id, nil
}

func ReindexProject(eclient *elastic.Client, projectID string, projectPage types.Project)error {
	//MODIFIES AN EXISTING PROJ (requires an elastic client pointer, string with project id, 
	//	and the modified project as a project type)
	//RETURNS AN ERROR
	ctx := context.Background()
	
	exists, err := eclient.IndexExists(PROJECT_INDEX).Do(ctx)

	if err != nil {return err }
	if !exists {return errors.New("Index does not exist")}

	_, err = eclient.Index().
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		Id(projectID).
		BodyJson(projectPage).
		Do(ctx)

	if err != nil {return err}

	return nil
}

func UpdateProject(eclient *elastic.Client, projectID string, field string, newContent interface{}) error{
	ctx:=context.Background()

	exists, err := eclient.IndexExists(PROJECT_INDEX).Do(ctx)
	if err != nil {return err}
	if !exists {return errors.New("Index does not exist")}

	_, err = get.GetProjectByID(eclient, projectID)
	if (err!=nil){return err}

	_, err = eclient.Update().
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		Id(projectID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	//if err != nil {return err}

	return nil
}



func AppendToProject(eclient *elastic.Client, projID string, field string, data interface{})error{return nil}//RETURN HERE
func RemoveFromProject(eclient *elastic.Client, projID string, field string, idx int, data interface{})error{return nil}




//Appends a new member
func AppendMember(eclient *elastic.Client, projectID string, member types.Member)error{
	ctx:= context.Background()
	proj, err := get.GetProjectByID(eclient, projectID)

	if (err!=nil) {return errors.New("Project does not exist")}

	
	 proj.Members = append(proj.Members, member)

	_,err =  eclient.Update().
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		Id(projectID).
		Doc(map[string]interface{}{"Members": proj.Members}).
		Do(ctx)

	return err
	
}


//Deletes a member
func DeleteMember(eclient *elastic.Client, projID string, memberID string,idx int)error{
	ctx:= context.Background()
	proj, err := get.GetProjectByID(eclient, projID)
	if (err!=nil) {return errors.New("Project does not exist")}
	
	
	proj.Members = append(proj.Members[:idx],proj.Members[idx+1:]...) //maintains order, while appending everything except the element at idx

	_,err =  eclient.Update().
		Index(USER_INDEX).
		Type(USER_TYPE).
		Id(projID).
		Doc(map[string]interface{}{"Members": proj.Members}).
		Do(ctx)
	
	return err
	

	
}

//Add QuickLink
func AppendProjectLink(eclient *elastic.Client, projectID string, link types.Link)error{
	ctx:= context.Background()
	proj, err := get.GetProjectByID(eclient, projectID)

	if (err!=nil) {return errors.New("Project does not exist")} //return error if nonexistent

	
	proj.QuickLinks = append(proj.QuickLinks,link) //retreive project

	_,err =  eclient.Update(). 
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		Id(projectID).
		Doc(map[string]interface{}{"QuickLinks": proj.QuickLinks}). //update project doc with new link appended to array
		Do(ctx)

	return err
	
}


//Delete QuickLink
func DeleteProjectLink(eclient *elastic.Client, projectID string, link types.Link,idx int)error{
	ctx:= context.Background()
	proj, err := get.GetProjectByID(eclient, projectID)
	if (err!=nil) {return errors.New("Project does not exist")}
	
	
	proj.QuickLinks = append(proj.QuickLinks[:idx],proj.QuickLinks[idx+1:]...) //append all elements BUT element at idx, maintains order

	_,err =  eclient.Update().
		Index(PROJECT_INDEX).
		Type(PROJECT_TYPE).
		Id(projectID).
		Doc(map[string]interface{}{"Quicklinks": proj.QuickLinks}). 
		Do(ctx)
	
	return err
	

	
}



