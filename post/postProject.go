package post

import (
	"context"
	"errors"
	"sync"

	get "github.com/sea350/ustart_go/get"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

const projectIndex = "test-project_data"
const projectType = "PROJECT"

var genericProjectUpdateLock sync.Mutex
var modifyMemberLock sync.Mutex

const projMapping = `
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

//IndexProject ... ADDS NEW PROJECT TO ES RECORDS
//Needs a type Project struct
//returns the new project's id and an error
func IndexProject(eclient *elastic.Client, newProj types.Project) (string, error) {

	tempString := "Project has not been indexed"
	ctx := context.Background()

	exists, err := eclient.IndexExists(projectIndex).Do(ctx)

	if err != nil {
		return tempString, err
	}

	if !exists {
		return tempString, errors.New("Index does not exist")
	}

	storedProj, err := eclient.Index().
		Index(projectIndex).
		Type(projectType).
		BodyJson(newProj).
		Do(ctx)

	if err != nil {
		return tempString, err
	}

	return storedProj.Id, nil
}

//ReindexProject ... REPLACES EXISTING ES DOC
//Specify the docid to be replaced and a type Project struct
//returns an error
func ReindexProject(eclient *elastic.Client, projectID string, projectPage types.Project) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(projectIndex).Do(ctx)

	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		BodyJson(projectPage).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}

//UpdateProject ... UPDATES A SINGLE FEILD IN AN EXISTING ES DOC
//Requires the docID, feild to be modified, and the new content
//Returns an error
func UpdateProject(eclient *elastic.Client, projectID string, field string, newContent interface{}) error {
	ctx := context.Background()

	exists, err := eclient.IndexExists(projectIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	genericProjectUpdateLock.Lock()
	defer genericProjectUpdateLock.Unlock()

	_, err = get.GetProjectByID(eclient, projectID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)
	//if err != nil {return err}

	return nil
}

//AppendMember ... APPENDS A NEW MEMBER TO AN EXISTING PROJECT DOC
//Requires the project docID and a type Member
//Returns an error
func AppendMember(eclient *elastic.Client, projectID string, member types.Member) error {

	ctx := context.Background()

	modifyMemberLock.Lock()
	defer modifyMemberLock.Unlock()

	proj, err := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	proj.Members = append(proj.Members, member)

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{"Members": proj.Members}).
		Do(ctx)

	return err

}

//DeleteMember ... REMOVES A SPECIFIC MEMBER FROM AN ARRAY
//Requires project docID and a type member
//Returns an error
func DeleteMember(eclient *elastic.Client, projID string, member types.Member) error {

	ctx := context.Background()

	modifyMemberLock.Lock()
	defer modifyMemberLock.Unlock()

	proj, err := get.GetProjectByID(eclient, projID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	index := -1
	for i := range proj.Members {
		if proj.Members[i] == member {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("Member not found")
	}

	proj.Members = append(proj.Members[:index], proj.Members[index+1:]...)

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projID).
		Doc(map[string]interface{}{"Members": proj.Members}).
		Do(ctx)

	return err

}

//AppendProjectLink ... ADDS A LINK TYPE TO A PROJECT'S QUICKLINKS
//Requires project's docID and a type link
//Returns an error
func AppendProjectLink(eclient *elastic.Client, projectID string, link types.Link) error {
	ctx := context.Background()

	proj, err := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	} //return error if nonexistent

	proj.QuickLinks = append(proj.QuickLinks, link) //retreive project

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{"QuickLinks": proj.QuickLinks}). //update project doc with new link appended to array
		Do(ctx)

	return err

}

//DeleteProjectLink ... ADDS A LINK TYPE TO A PROJECT'S QUICKLINKS
//Requires project's docID and a type link
//Returns an error
func DeleteProjectLink(eclient *elastic.Client, projectID string, link types.Link) error {
	ctx := context.Background()
	proj, err := get.GetProjectByID(eclient, projectID)
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
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{"Quicklinks": proj.QuickLinks}).
		Do(ctx)

	return err

}

//AppendMemberReqSent ... APPENDS A USER ID TO A PROJECTS MemberReqSent ARRAY
//Requires project's docID and the user's docID
//Returns an error
func AppendMemberReqSent(eclient *elastic.Client, projectID string, userID string) error {

	ctx := context.Background()

	proj, err := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	} //return error if nonexistent

	proj.MemberReqSent = append(proj.MemberReqSent, userID) //retreive project

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{"MemberReqSent": proj.MemberReqSent}). //update project doc with new link appended to array
		Do(ctx)

	return err
}

//DeleteMemberReqSent ... DELETES A USER ID IN A PROJECT'S MemberReqSent ARRAY
//Requires project's docID and the user's docID
//Returns an error
func DeleteMemberReqSent(eclient *elastic.Client, projectID string, userID string) error {
	ctx := context.Background()
	proj, err := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	//replace with universal.FindIndex when it works
	index := -1
	for i := range proj.MemberReqSent {
		if proj.MemberReqSent[i] == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("link not found")
	}

	proj.MemberReqSent = append(proj.MemberReqSent[:index], proj.MemberReqSent[index+1:]...)

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{"MemberReqSent": proj.MemberReqSent}).
		Do(ctx)

	return err
}

//AppendMemberReqReceived ... APPENDS A USER ID TO A PROJECTS MemberReqReceived ARRAY
//Requires project's docID and the user's docID
//Returns an error
func AppendMemberReqReceived(eclient *elastic.Client, projectID string, userID string) error {

	ctx := context.Background()

	proj, err := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	} //return error if nonexistent

	proj.MemberReqReceived = append(proj.MemberReqReceived, userID) //retreive project

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{"MemberReqReceived": proj.MemberReqReceived}). //update project doc with new link appended to array
		Do(ctx)

	return err
}

//DeleteMemberReqReceived ... DELETES A USER ID IN A PROJECT'S MemberReqReceived ARRAY
//Requires project's docID and the user's docID
//Returns an error
func DeleteMemberReqReceived(eclient *elastic.Client, projectID string, userID string) error {
	ctx := context.Background()
	proj, err := get.GetProjectByID(eclient, projectID)
	if err != nil {
		return errors.New("Project does not exist")
	}

	//replace with universal.FindIndex when it works
	index := -1
	for i := range proj.MemberReqReceived {
		if proj.MemberReqReceived[i] == userID {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("link not found")
	}

	proj.MemberReqReceived = append(proj.MemberReqReceived[:index], proj.MemberReqReceived[index+1:]...)

	_, err = eclient.Update().
		Index(projectIndex).
		Type(projectType).
		Id(projectID).
		Doc(map[string]interface{}{"MemberReqReceived": proj.MemberReqReceived}).
		Do(ctx)

	return err

}
