package uses

const entryIndex = "test-entry_data"
const entryType = "ENTRY"

//THIS FILE IS RETIRED

/*
func DeleteLike(eclient *elastic.Client, entryID string, likerID string) error {
	ctx := context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)
	var idx int

	for i := range anEntry.Likes {
		if likerID == anEntry.Likes[i].UserID {
			idx = i
		}
	}

	anEntry.Likes = append(anEntry.Likes[:idx], anEntry.Likes[idx+1:]...)

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)

	return err

}

func AppendLike(eclient *elastic.Client, entryID string, likerID string) error {
	ctx := context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)
	newLike := types.Like{}
	newLike.UserID = likerID
	//newLike.TimeStamp = time.Now()
	anEntry.Likes = append(anEntry.Likes, newLike)
	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"Likes": anEntry.Likes}).
		Do(ctx)

	return err

}

func CheckLike(eclient *elastic.Client, entryID string, theLike types.Like, action bool) error {
	isAppended := false

	for isAppended == false {
		theDoc, err := get.GetEntryByID(eclient, entryID)
		if err != nil {
			return errors.New("Entry does not exist")
		}

		for i := range theDoc.Likes {

			if theDoc.Likes[i] == theLike {
				if action == true {
					isAppended = true
					return nil

				} else {
					isAppended = false
				}
			}
		}

		if action == true && isAppended == false {
			checkErr := AppendLike(eclient, entryID, theLike.UserID)
			if checkErr != nil {
				return checkErr
			}

		} else if action == false && isAppended == false {
			return nil
		}

	}

	if action == false && isAppended == true {
		checkErr := DeleteLike(eclient, entryID, theLike.UserID)
		if checkErr != nil {
			return checkErr
		}
	}

	return nil
}

func AppendShareID(eclient *elastic.Client, entryID string, shareID string, idx int) error {
	ctx := context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)

	anEntry.ShareIDs = append(anEntry.ShareIDs, shareID)

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	return err

}

func DeleteShareID(eclient *elastic.Client, entryID string, shareID string, idx int) error {
	ctx := context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)

	for i := range anEntry.ShareIDs {
		if shareID == anEntry.ShareIDs[i] {
			idx = i
		}
	}

	anEntry.ShareIDs = append(anEntry.ShareIDs[:idx], anEntry.ShareIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ShareIDs": anEntry.ShareIDs}).
		Do(ctx)

	return err

}

func CheckShareID(eclient *elastic.Client, entryID string, shareID string, action bool, idx int) error {
	isAppended := false

	for isAppended == false {
		theDoc, err := get.GetEntryByID(eclient, entryID)
		if err != nil {
			return errors.New("Entry does not exist")
		}

		for i := range theDoc.ShareIDs {

			if theDoc.ShareIDs[i] == shareID {
				if action == true {
					isAppended = true
					return nil

				} else {
					isAppended = false
				}
			}
		}

		if action == true && isAppended == false {
			checkErr := AppendShareID(eclient, entryID, shareID, idx)
			if checkErr != nil {
				return checkErr
			}

		} else if action == false && isAppended == false {
			return nil
		}

	}

	if action == false && isAppended == true {
		checkErr := DeleteShareID(eclient, entryID, shareID, idx)
		if checkErr != nil {
			return checkErr
		}
	}

	return nil

}

func DeleteReplyID(eclient *elastic.Client, entryID string, replyID string, idx int) error {
	ctx := context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)

	for i := range anEntry.ShareIDs {
		if replyID == anEntry.ReplyIDs[i] {
			idx = i
		}
	}

	anEntry.ReplyIDs = append(anEntry.ReplyIDs[:idx], anEntry.ReplyIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": anEntry.ReplyIDs}).
		Do(ctx)

	return err

}

func AppendReplyID(eclient *elastic.Client, entryID string, replyID string) error {
	//newLike.TimeStamp = time.Now()
	ctx := context.Background()
	anEntry, err := get.GetEntryByID(eclient, entryID)
	anEntry.ReplyIDs = append(anEntry.ReplyIDs, replyID)
	_, err = eclient.Update().
		Index(entryIndex).
		Type(entryType).
		Id(entryID).
		Doc(map[string]interface{}{"ReplyIDs": anEntry.ReplyIDs}).
		Do(ctx)

	return err

}

func CheckReplyID(eclient *elastic.Client, entryID string, replyID string, action bool, idx int) error {
	isAppended := false

	for isAppended == false {
		theDoc, err := get.GetEntryByID(eclient, entryID)
		if err != nil {
			return errors.New("Entry does not exist")
		}

		for i := range theDoc.ReplyIDs {

			if theDoc.ReplyIDs[i] == replyID {
				if action == true {
					isAppended = true
					return nil

				} else {
					isAppended = false
				}
			}
		}

		if action == true && isAppended == false {
			checkErr := AppendReplyID(eclient, entryID, replyID)
			if checkErr != nil {
				return checkErr
			}

		} else if action == false && isAppended == false {
			return nil
		}

	}

	if action == false && isAppended == true {
		checkErr := DeleteReplyID(eclient, entryID, replyID, idx)
		if checkErr != nil {
			return checkErr
		}
	}

	return nil

}
*/
