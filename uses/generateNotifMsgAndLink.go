package uses

import (
	"errors"
	"strconv"

	"github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//GenerateNotifMsgAndLink ... creates the text representation of a notification
// Return message, link, and error
func GenerateNotifMsgAndLink(eclient *elastic.Client, notif types.Notification) (string, string, error) {
	switch notif.Class {
	case 1:
		if len(notif.ReferenceIDs) > 1 {
			return strconv.Itoa(len(notif.ReferenceIDs)) + " new likes on your journal entry", `/404`, nil
		} else if len(notif.ReferenceIDs) == 1 {
			head, err := ConvertUserToFloatingHead(eclient, notif.ReferenceIDs[0])
			if err != nil {
				return ``, ``, err
			}
			return head.FirstName + ` ` + head.LastName + ` has liked your journal entry`, `/404`, nil
		} else {
			return ``, ``, errors.New("invalid notification")
		}
	case 2:
		if len(notif.ReferenceIDs) > 1 {
			return strconv.Itoa(len(notif.ReferenceIDs)) + " new comments on your journal entry", `/404`, nil
		} else if len(notif.ReferenceIDs) == 1 {
			head, err := ConvertUserToFloatingHead(eclient, notif.ReferenceIDs[0])
			if err != nil {
				return ``, ``, err
			}
			return head.FirstName + ` ` + head.LastName + ` has commented to your journal entry`, `/404`, nil
		} else {
			return ``, ``, errors.New("invalid notification")
		}
	case 3:
		if len(notif.ReferenceIDs) > 1 {
			return strconv.Itoa(len(notif.ReferenceIDs)) + " people have shared your journal entry", `/404`, nil
		} else if len(notif.ReferenceIDs) == 1 {
			head, err := ConvertUserToFloatingHead(eclient, notif.ReferenceIDs[0])
			if err != nil {
				return ``, ``, err
			}
			return head.FirstName + ` ` + head.LastName + ` has shared your journal entry`, `/404`, nil
		} else {
			return ``, ``, errors.New("invalid notification")
		}
	case 4:
		if len(notif.ReferenceIDs) > 1 {
			return "You have " + strconv.Itoa(len(notif.ReferenceIDs)) + " new followers", `/404`, nil
		} else if len(notif.ReferenceIDs) == 1 {
			head, err := ConvertUserToFloatingHead(eclient, notif.ReferenceIDs[0])
			if err != nil {
				return ``, ``, err
			}
			return head.FirstName + ` ` + head.LastName + ` is now following you`, `/profile/` + head.Username, nil
		} else {
			return ``, ``, errors.New("invalid notification")
		}
	case 11:
		head, err := ConvertUserToFloatingHead(eclient, notif.RedirectToID)
		if err != nil {
			return ``, ``, err
		}
		text := head.FirstName + ` ` + head.LastName + ` has requested to join `
		link := `/profile/` + head.Username
		head, err = ConvertProjectToFloatingHead(eclient, notif.ReferenceIDs[0])
		if err != nil {
			return ``, ``, err
		}
		text = text + head.FirstName
		return text, link, nil
	case 12:
		head, err := ConvertProjectToFloatingHead(eclient, notif.RedirectToID)
		if err != nil {
			return ``, ``, err
		}
		return head.FirstName + ` has accepted your request to join!`, `/Projects/` + head.Username, nil
	case 13:
		head, err := ConvertProjectToFloatingHead(eclient, notif.RedirectToID)
		if err != nil {
			return ``, ``, err
		}
		return head.FirstName + ` has accepted your request to join!`, `/Projects/` + head.Username, nil
	}

	return "", ``, errors.New("invalid notif")
}
