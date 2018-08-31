package uses

import (
	getChat "github.com/sea350/ustart_go/get/chat"
	getEvent "github.com/sea350/ustart_go/get/event"
	getProject "github.com/sea350/ustart_go/get/project"
	getUser "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ConvertUserToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertUserToFloatingHead(eclient *elastic.Client, userDocID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	usr, err := getUser.UserByID(eclient, userDocID)
	if err != nil {
		return head, err
	}

	head.FirstName = usr.FirstName
	head.LastName = usr.LastName
	head.Image = usr.Avatar
	head.Username = usr.Username
	head.DocID = userDocID
	head.Email = usr.Email
	head.Tags = usr.Tags

	// head.Interface = usr.Tags
	head.Bio = usr.Description
	head.Category = usr.Majors[0]

	return head, err
}

//ConvertProjectToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertProjectToFloatingHead(eclient *elastic.Client, projectID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	proj, err := getProject.ProjectByID(eclient, projectID)
	if err != nil {
		panic(err)
	}

	head.FirstName = proj.Name
	head.Bio = proj.Description
	head.Image = proj.Avatar
	head.Username = proj.URLName
	head.DocID = projectID
	head.Notifications = len(proj.MemberReqReceived)
	head.Interface = proj.Tags
	head.Tags = proj.Tags
	head.Category = proj.Category

	return head, err
}

//ConvertEventToFloatingHead ... pulls latest version of user and converts relevent data into floating head
func ConvertEventToFloatingHead(eclient *elastic.Client, eventID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	evnt, err := getEvent.EventByID(eclient, eventID)
	if err != nil {
		panic(err)
	}

	head.FirstName = evnt.Name
	head.Bio = evnt.Description
	head.Image = evnt.Avatar
	head.Username = evnt.URLName
	head.DocID = eventID
	head.Notifications = len(evnt.MemberReqReceived)
	head.Interface = evnt.Tags
	head.Category = evnt.Category

	return head, err
}

//ConvertChatToFloatingHead ... pulls latest version of chat and converts relevent data into floating head
func ConvertChatToFloatingHead(eclient *elastic.Client, conversationID string, viewerID string) (types.FloatingHead, error) {
	var head types.FloatingHead

	convo, err := getChat.ConvoByID(eclient, conversationID)
	if err != nil {
		return head, err
	}

	if convo.Class == 3 {
		head, err = ConvertProjectToFloatingHead(eclient, convo.ReferenceID)
		if err != nil {
			return head, err
		}
		head.Classification = 3
	}

	if convo.Class == 4 {
		head, err = ConvertEventToFloatingHead(eclient, convo.ReferenceID)
		if err != nil {
			return head, err
		}
	}

	var msg types.Message
	if len(convo.MessageIDArchive) > 0 {
		msg, err = getChat.MsgByID(eclient, convo.MessageIDArchive[len(convo.MessageIDArchive)-1])
		if err != nil {
			return head, err
		}
		head.Time = msg.TimeStamp
	}

	for i := range convo.Eavesdroppers {
		if convo.Eavesdroppers[i].DocID == viewerID {
			head.Notifications = len(convo.MessageIDArchive) - convo.Eavesdroppers[i].Bookmark - 1
			if len(convo.Eavesdroppers) == 1 && convo.Class == 1 {
				usr, err := getUser.UserByID(eclient, convo.Eavesdroppers[i].DocID)
				if err != nil {
					return head, err
				}
				head.Username = usr.Username
				head.FirstName = usr.FirstName
				head.LastName = usr.LastName
				head.Image = usr.Avatar
			}
		} else if convo.Class == 1 {
			usr, err := getUser.UserByID(eclient, convo.Eavesdroppers[i].DocID)
			if err != nil {
				return head, err
			}
			head.Username = usr.Username
			head.FirstName = usr.FirstName
			head.LastName = usr.LastName
			head.Image = usr.Avatar
			if msg.SenderID == convo.Eavesdroppers[i].DocID {
				head.Bio = []rune(usr.FirstName + `: ` + msg.Content)
			}
		} else if msg.SenderID == convo.Eavesdroppers[i].DocID {
			usr, err := getUser.UserByID(eclient, convo.Eavesdroppers[i].DocID)
			if err != nil {
				return head, err
			}
			head.Bio = []rune(usr.FirstName + `: ` + msg.Content)
		}

	}

	if msg.SenderID == viewerID {
		head.Bio = []rune(`You: ` + msg.Content)
	}
	head.DocID = conversationID

	return head, err
}
