package uses

import (
	"errors"
	"strconv"

	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//GenerateMessage ...
func GenerateMessage(eclient *elastic.Client, notif types.Notification) (string, error) {
	switch notif.Class {
	case 4:
		if len(notif.ReferenceIDs) > 1 {
			return strconv.Itoa(len(notif.ReferenceIDs)) + " new followers", nil
		} else if len(notif.ReferenceIDs) == 1 {
			head, err := ConvertUserToFloatingHead(eclient, notif.ReferenceIDs[0])
			if err != nil {
				return ``, err
			}
			return head.FirstName + ` ` + head.LastName + ` is now following you`, nil
		} else {
			return ``, errors.New("invalid notification")
		}
	}

	return ``, errors.New("invalid notif")
}
