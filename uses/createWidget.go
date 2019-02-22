package uses

import (
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
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
