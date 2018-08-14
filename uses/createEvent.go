package uses

import (
	"errors"
	"strings"

	"time"

	eventGet "github.com/sea350/ustart_go/get/event"
	eventPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//CreateEvent ... Create an event
//Requires all fundemental information for the new event (title, creator docID, etc...)
//Returns an error if there was a problem with database submission
func CreateEvent(eclient *elastic.Client, title string, description []rune, makerID string, category string, location types.LocStruct, eventTimeStart time.Time, eventTimeEnd time.Time, college string, customURL string) (string, error) {
	inUse, err := eventGet.EventURLInUse(eclient, customURL)
	if err != nil {
		return "", err
	}
	if inUse {
		return "", errors.New("URL is taken")
	}

	var newEvent types.Events
	newEvent.Visible = true
	newEvent.CreationDate = time.Now()
	newEvent.Avatar = "https://i.imgur.com/TYFKsdi.png"
	newEvent.Name = title
	newEvent.Description = description
	newEvent.Host = makerID
	newEvent.Category = category
	newEvent.Location = location
	newEvent.EventDateStart = eventTimeStart
	newEvent.EventDateEnd = eventTimeEnd
	newEvent.Organization = college
	if college == `` {
		newEvent.Organization = ""
	}
	newEvent.URLName = customURL
	if customURL != `` {
		newEvent.URLName = customURL
	}

	var maker types.EventMembers
	maker.MemberID = makerID
	maker.Role = 0
	maker.Title = "Creator"
	maker.JoinDate = time.Now()
	maker.Visible = true

	newEvent.Members = append(newEvent.Members, maker)
	newEvent.PrivilegeProfiles = append(newEvent.PrivilegeProfiles, SetEventMemberPrivileges(0), SetEventMemberPrivileges(1), SetEventMemberPrivileges(2))

	id, err := eventPost.IndexEvent(eclient, newEvent)
	if err != nil {
		return id, err
	}

	var addEvnt types.EventInfo
	addEvnt.EventID = id
	addEvnt.Visible = true
	err = userPost.AppendEvent(eclient, makerID, addEvnt)
	if err != nil {
		panic(err)
		//return id, err
	}

	if customURL == `` {
		id = strings.ToLower(id)
		err = eventPost.UpdateEvent(eclient, id, "URLName", id)
	} else {
		id = customURL
	}
	return id, err
}
