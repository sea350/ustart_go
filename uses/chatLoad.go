package uses

import (
	"errors"

	getChat "github.com/sea350/ustart_go/get/chat"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ChatLoad ... agreggates a quantity of messages from a certain index in a conversation
//if you start from zero it pulls from the end of the array
//returns current size of messages array and err
//NOTE: this method may be replaced with a scrollservice elastic search
func ChatLoad(eclient *elastic.Client, convoID string, startFrom int, pullAmount int) (int, []types.Message, error) {
	chat, err := getChat.ConvoByID(eclient, convoID)
	if err != nil {
		return 0, []types.Message{}, err
	}

	if startFrom < 0 || pullAmount < 0 {
		return 0, []types.Message{}, errors.New("Out of bounds")
	}

	length := len(chat.MessageIDArchive)

	if startFrom > length {
		startFrom = length - 1
	}
	if startFrom == 0 {
		startFrom = length - 1
	}
	if startFrom-pullAmount < 0 {
		pullAmount = length
	}

	var problemMsgIDs string
	var messages []types.Message
	for i := startFrom; i > startFrom-pullAmount; i-- {
		message, err := getChat.MsgByID(eclient, chat.MessageIDArchive[i])
		if err != nil {
			problemMsgIDs = problemMsgIDs + ", " + chat.MessageIDArchive[i]
			continue
		}
		messages = append(messages, message)
	}
	if len(problemMsgIDs) > 0 {
		return startFrom - pullAmount, messages, errors.New("Problems loading the following messages" + problemMsgIDs)
	}
	return startFrom - pullAmount, messages, nil
}
