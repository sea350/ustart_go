package uses

import (
	"strings"

	"errors"
	"time"

	projGet "github.com/sea350/ustart_go/get/project"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//CreateProject ... CREATE A NORMAL PROJECT
//Requires all fundamental information for the new project (title, creator docID, etc ...)
//Returns an error if there was a problem with database submission
func CreateProject(eclient *elastic.Client, title string, description []rune, makerID string, category string, college string, customURL string) (string, error) {
	inUse, err := projGet.URLInUse(eclient, customURL)
	if err != nil {
		return "", err
	}
	if inUse {
		return "", errors.New("URL is taken")
	}

	var newProj types.Project
	newProj.Name = title
	newProj.Description = description
	newProj.Visible = true
	newProj.CreationDate = time.Now()
	newProj.Avatar = "https://i.imgur.com/TYFKsdi.png"
	newProj.Category = category
	newProj.Organization = college

	var maker types.Member
	maker.JoinDate = time.Now()
	maker.MemberID = makerID
	maker.Role = 0
	maker.Title = "Creator"
	maker.Visible = true

	newProj.Members = append(newProj.Members, maker)
	newProj.PrivilegeProfiles = append(newProj.PrivilegeProfiles, SetMemberPrivileges(0), SetMemberPrivileges(1), SetMemberPrivileges(2))

	id, err := projPost.IndexProject(eclient, newProj)
	if err != nil {
		return id, err
	}
	var addProj types.ProjectInfo
	addProj.ProjectID = id
	addProj.Visible = true
	err = userPost.AppendProject(eclient, makerID, addProj)
	if err != nil {
		panic(err)
	}

	if customURL == `` {
		err = projPost.UpdateProject(eclient, id, "URLName", strings.ToLower(id))
		id = strings.ToLower(id)
	} else {
		err = projPost.UpdateProject(eclient, id, "URLName", customURL)
		id = customURL
	}

	return id, err

}
