package get

import (
	"context"
	"encoding/json"
	"errors"

	elastic "github.com/olivere/elastic"
	get "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
)

//DMExists ...
//checks to see if a conversation between 2 people already exists
func DMExists(eclient *elastic.Client, eavesdropperOne string, eavesdropperTwo string) (bool, string, error) {

	eavesOne, errOne := get.UserExists(eclient, eavesdropperOne)
	if !eavesOne {
		return false, "", errors.New("E1: Not all participants exist")
	}

	if errOne != nil {
		return false, "", errOne
	}

	eavesTwo, errTwo := get.UserExists(eclient, eavesdropperTwo)
	if !eavesTwo {
		return false, "", errors.New("E2: Not all participants exist")
	}
	if errTwo != nil {
		return false, "", errTwo
	}

	// var dash = rune('-')
	// var underscore = rune('_')
	// var tempRuneArr []rune
	// for _, char := range eavesdropperOne {
	// 	if char != dash && char != underscore {
	// 		tempRuneArr = append(tempRuneArr, char)
	// 	}
	// }
	// trimmedID1 := string(tempRuneArr)

	// tempRuneArr = []rune{}
	// for _, char := range eavesdropperTwo {
	// 	if char != dash && char != underscore {
	// 		tempRuneArr = append(tempRuneArr, char)
	// 	}
	// }
	// trimmedID2 := string(tempRuneArr)

	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// log.Println("debug text 1 ", trimmedID1)
	// log.Println("debug text 2 ", trimmedID2)
	// for eavesdropperOne[0] == dash || eavesdropperOne[0] == underscore {
	// 	eavesdropperOne = eavesdropperOne[1:]
	// }
	// for eavesdropperOne[len(eavesdropperOne)-1] == dash || eavesdropperOne[len(eavesdropperOne)-1] == underscore {
	// 	eavesdropperOne = eavesdropperOne[:len(eavesdropperOne)-1]
	// }

	// for eavesdropperTwo[0] == dash || eavesdropperTwo[0] == underscore {
	// 	eavesdropperTwo = eavesdropperTwo[1:]
	// }
	// for eavesdropperTwo[len(eavesdropperTwo)-1] == dash || eavesdropperTwo[len(eavesdropperTwo)-1] == underscore {
	// 	eavesdropperTwo = eavesdropperTwo[:len(eavesdropperTwo)-1]
	// }

	query := elastic.NewBoolQuery()

	query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID.keyword", eavesdropperOne))
	query = query.Must(elastic.NewTermQuery("Eavesdroppers.DocID.keyword", eavesdropperTwo))
	query = query.Must(elastic.NewTermQuery("Class", "1"))

	if eavesdropperOne == eavesdropperTwo {
		query = query.Must(elastic.NewTermQuery("Size", "1"))
	} else {
		query = query.Must(elastic.NewTermQuery("Size", "2"))
	}

	ctx := context.Background() //intialize context background
	searchResults, err := eclient.Search().
		Index(globals.ConvoIndex).
		Query(query).
		Pretty(true).
		Do(ctx)

	var chatID string

	if err != nil {
		return false, chatID, err
	}

	exists := searchResults.TotalHits() != 0
	if !exists {
		return exists, chatID, err

	}
	// multi := searchResults.TotalHits() > 1
	// if multi {
	// 	return exists, chatID, errors.New("Too many chats of this type exist")
	// }

	var convo types.Conversation
	for _, ch := range searchResults.Hits.Hits {
		err := json.Unmarshal(*ch.Source, &convo) //unmarshal type RawMessage into user struct
		if err != nil {
			return false, chatID, err
		}
		var thankYouNext bool
		for _, eaves := range convo.Eavesdroppers {
			if eaves.DocID != eavesdropperOne && eaves.DocID != eavesdropperTwo {
				thankYouNext = true
				break
			}
		}
		if thankYouNext {
			continue
		}
		chatID = ch.Id
		return exists, chatID, err
	}

	return false, chatID, err

}
