package uses

import (
	"strings"

	"time"

	eventPost "github.com/sea350/ustart_go/post/event"
	userPost "github.com/sea350/ustart_go/post/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//CreateEvent ... Create an event
//Requires all fundemental information for the new event (title, creator docID, etc...)
//Returns an error if there was a problem with database submission
func CreateEvent(eclient *elastic.Client, title string, description []rune, makerID string, category string, location types.LocStruct, eventTimeStart time.Time) (string, error) {
	var newEvent types.Events
	newEvent.Name = title
	newEvent.Description = description
	newEvent.Visible = true
	newEvent.CreationDate = time.Now()
	newEvent.Avatar = "https://i.imgur.com/TYFKsdi.png"
	newEvent.EventDateStart = eventTimeStart
	newEvent.EventDateEnd = eventTimeEnd
	newEvent.Location = location
	newEvent.Category = category
	newEvent.Host = makerID

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
	}

	err = eventPost.UpdateEvent(eclient, id, "URLName", strings.ToLower(id))
	id = strings.ToLower(id)

	return id, err
}
