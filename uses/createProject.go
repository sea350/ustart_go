package uses

import (
	"strings"

	"errors"
	"time"

	getChat "github.com/sea350/ustart_go/get/chat"
	projGet "github.com/sea350/ustart_go/get/project"
	postChat "github.com/sea350/ustart_go/post/chat"
	followPost "github.com/sea350/ustart_go/post/follow"
	projPost "github.com/sea350/ustart_go/post/project"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//CreateProject ... CREATE A NORMAL PROJECT
//Requires all fundamental information for the new project (title, creator docID, etc ...)
//Returns an error if there was a problem with database submission
func CreateProject(eclient *elastic.Client, title string, description []rune, makerID string, category string, college string, customURL string, location types.LocStruct) (string, error) {
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
	newProj.Avatar = "https://i.imgur.com/8BnkFLO.png"
	newProj.Banner = "https://i.imgur.com/XTj1t1J.png"
	newProj.Category = category
	newProj.Organization = college
	newProj.Location = location
	if customURL != `` {
		newProj.URLName = customURL
	}

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
		return id, err
	}

	errFollow := followPost.IndexFollow(eclient, id)
	if errFollow != nil {
		return id, errFollow
	}
	var newConvo types.Conversation
	newConvo.Class = 3
	newConvo.Title = "General"
	newConvo.ReferenceID = id

	newConvo.Eavesdroppers = append(newConvo.Eavesdroppers, types.Eavesdropper{Class: 1, DocID: makerID})

	convoID, err := postChat.IndexConvo(eclient, newConvo)
	if err != nil {
		return id, err
	}

	proxyID, err := getChat.ProxyIDByUserID(eclient, makerID)
	if err != nil {
		return id, err
	}

	err = postChat.AppendToProxy(eclient, proxyID, convoID, false)
	if err != nil {
		return id, err
	}

	err = projPost.UpdateProject(eclient, id, "Subchats", []types.Subchat{types.Subchat{ConversationID: convoID, ChatName: "General Chat"}})
	if err != nil {
		return id, err
	}

	if customURL == `` {
		id = strings.ToLower(id)
		err = projPost.UpdateProject(eclient, id, "URLName", id)
	} else {
		id = customURL
	}

	var initMsg types.Message
	initMsg.SenderID = makerID
	initMsg.ConversationID = convoID
	initMsg.TimeStamp = time.Now()
	initMsg.Content = "Welcome to your new project!"
	initMsg.Hidden = false

	err = ChatSend(eclient, initMsg)
	if err != nil {
		return id, err
	}

	return id, err

}
