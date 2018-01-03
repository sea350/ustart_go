package uses

import (
	"sync"
	"time"

	post "github.com/sea350/ustart_go/post/project"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

var memberModLock sync.Mutex

//CreateProject ... CREATE A NORMAL PROJECT
//Requires all fundamental information for the new project (title, creator docID, etc ...)
//Returns an error if there was a problem with database submission
func CreateProject(eclient *elastic.Client, title string, description []rune, makerID string) error {
	var newProj types.Project
	newProj.Name = title
	newProj.Description = description
	newProj.Visible = true
	newProj.CreationDate = time.Now()

	var maker types.Member
	maker.JoinDate = time.Now()
	maker.MemberID = makerID
	maker.Role = 0
	maker.Title = "Creator"
	maker.Visible = true

	newProj.Members = append(newProj.Members, maker)

	id, err := post.IndexProject(eclient, newProj)
	if err != nil {
		return err
	}
	var addProj types.ProjectInfo
	addProj.ProjectID = id
	addProj.Visible = true
	err = post.AppendProject(eclient, makerID, addProj)

	return err

}
