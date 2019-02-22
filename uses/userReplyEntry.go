package uses

import (
	"time"

	postEntry "github.com/sea350/ustart_go/post/entry"
	types "github.com/sea350/ustart_go/types"
	elastic "github.com/olivere/elastic"
)

//UserReplyEntry ... CREATES A REPLY ENTRY FROM A USER
//Requires the user's docID, the parent entry docID and the content of the post
//Returns an error
func UserReplyEntry(eclient *elastic.Client, userID string, entryID string, content []rune) error {

	var newReply types.Entry
	newReply.PosterID = userID
	newReply.Content = content
	newReply.ReferenceEntry = entryID
	newReply.TimeStamp = time.Now()
	newReply.Classification = 1
	newReply.Visible = true

	replyID, err := postEntry.IndexEntry(eclient, newReply)
	if err != nil {
		return err
	}

	/*err = postUser.AppendEntryID(eclient, userID, replyID)
	if err != nil {
		return err
	}*/

	err = postEntry.AppendReplyID(eclient, entryID, replyID)
	return err
}
