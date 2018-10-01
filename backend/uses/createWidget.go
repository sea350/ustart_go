package uses

import (
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//CreateWidget ...
func CreateWidget(eclient *elastic.Client, userID string, newLink string, newPos int, newClass int) types.Widget {
	var newWidget types.Widget

	newWidget.UserID = userID
	//newWidget.Link = newLink
	newWidget.Position = newPos
	newWidget.Classification = newClass

	return newWidget
}
