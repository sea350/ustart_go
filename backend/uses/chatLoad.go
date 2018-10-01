package uses

import (
	"errors"

	getChat "github.com/sea350/ustart_go/backend/get/chat"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatLoad ... agreggates a quantity of messages from a certain index in a conversation
//if you start from -1 it pulls from the end of the array
//returns current size of messages array and err
//NOTE: this method may be replaced with a scrollservice elastic search
func ChatLoad(eclient *elastic.Client, convoID string, startFrom int, pullAmount int) (int, []types.Message, error) {
	chat, err := getChat.ConvoByID(eclient, convoID)
	if err != nil {
		return 0, []types.Message{}, err
	}

	if startFrom < -1 || pullAmount < 0 || startFrom == 0 {
		return 0, []types.Message{}, errors.New("Out of bounds")
	}

	length := len(chat.MessageIDArchive)

	if startFrom > length {
		startFrom = length - 1
	}
	if startFrom == -1 {
		startFrom = length - 1
	}
	if startFrom-pullAmount < 0 {
		pullAmount = startFrom
	}

	//NOTE, the pull is inverted against terminology
	//startFrom acts as an end at
	//startfrom-pullamount is the start
	var problemMsgIDs string
	var messages []types.Message
	for i := startFrom - pullAmount; i <= startFrom; i++ {
		message, err := getChat.MsgByID(eclient, chat.MessageIDArchive[i])
		if err != nil {
			problemMsgIDs = problemMsgIDs + ", " + chat.MessageIDArchive[i]
			continue
		}
		messages = append(messages, message)
	}
	final := startFrom - pullAmount
	if len(problemMsgIDs) > 0 {
		return final, messages, errors.New("Problems loading the following messages" + problemMsgIDs)
	}
	return final, messages, nil
}
