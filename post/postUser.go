package post

import (
	"context"
	"errors"
	"sync"

	get "github.com/sea350/ustart_go/get"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

const esUserIndex = "test-user_data"
const esUserType = "USER"

const mapping = `
{
    "mappings":{
        "User":{
            "properties":{
                "Email":{
                    "type":"keyword"
                },
                "Username":{
                	"type":"keyword"
                },
                "AccCreation":{
                	"type": date"
				},
				"FirstName":{
					"type": "keyword"
				},
				"LastName":{
					"type":"keyword"
				}

                
            }
        }
    }
}`

var procLock sync.Mutex
var likeLock sync.Mutex
var colleagueLock sync.Mutex
var followLock sync.Mutex
var projectLock sync.Mutex
var blockLock sync.Mutex
var tagLock sync.Mutex
var entryLock sync.Mutex

/*TODO: Make this function much better*/

//IndexUser ...
// adds a new user document to the ES cluster
// returns err, nil if successful.
func IndexUser(eclient *elastic.Client, newAcc types.User) error {
	// Check if the index exists
	ctx := context.Background()
	exists, err := eclient.IndexExists(esUserIndex).Do(ctx)
	if err != nil {
		return err
	}
	// If the index doesn't exist, create it and return error.
	if !exists {
		createIndex, Err := eclient.CreateIndex(esUserIndex).BodyString(mapping).Do(ctx)
		if Err != nil {
			_, _ = eclient.IndexExists(esUserIndex).Do(ctx)
			panic(Err)
		}
		// TODO fix this.
		if !createIndex.Acknowledged {
		}

		// Return an error saying it doesn't exist
		return errors.New("Index does not exist")
	}

	// Index the document.
	_, Err := eclient.Index().
		Index(esUserIndex).
		Type(esUserType).
		BodyJson(newAcc).
		Do(ctx)

	if Err != nil {
		return Err
	}

	return nil
}

//ReindexUser ...
//  Add a new user to ES.
//  Returns an error, nil if successful
func ReindexUser(eclient *elastic.Client, userID string, userAcc types.User) error {

	ctx := context.Background()
	exists, err := eclient.IndexExists(esUserIndex).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = eclient.Index().
		Index(esUserIndex).
		Type(esUserType).
		Id(userID).
		BodyJson(userAcc).
		Do(ctx)

	if err != nil {
		return err
	}

	return nil
}

//UpdateUser ...
//  Change a single field of the ES Document
//  Return an error, nil if successful
func UpdateUser(eclient *elastic.Client, userID string, field string, newContent interface{}) error {

	ctx := context.Background()

	exists, err := eclient.IndexExists(esUserIndex).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("Index does not exist")
	}

	_, err = get.GetUserByID(eclient, userID)
	if err != nil {
		return err
	}

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(userID).
		Doc(map[string]interface{}{field: newContent}).
		Do(ctx)

	return err
}

//AppendCollReq ...
//  Appends to either sent or received collegue, based on whichOne
//  True = sent; False = received.
func AppendCollReq(eclient *elastic.Client, usrID string, collegueID string, whichOne bool) error {

	ctx := context.Background()

	colleagueLock.Lock()
	defer colleagueLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.SentCollReq = append(usr.SentCollReq, collegueID)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentCollReq": usr.SentCollReq}).
			Do(ctx)

		return err
	}
	usr.ReceivedCollReq = append(usr.ReceivedCollReq, collegueID)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedCollReq": usr.ReceivedCollReq}).
		Do(ctx)

	return err
}

//DeleteCollReq ...
//  Deletes from sent or received collegue request arrays depending on whichOne
//  True = sent; false = received
func DeleteCollReq(eclient *elastic.Client, usrID string, reqID string, whichOne bool) error {
	ctx := context.Background()

	followLock.Lock()
	defer followLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		index := -1
		for i := range usr.SentCollReq {
			if usr.SentCollReq[i] == reqID {
				index = i
			}
		}
		if index < 0 {
			return errors.New("Index not found")
		}

		usr.SentCollReq = append(usr.SentCollReq[:index], usr.SentCollReq[index+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentCollReq": usr.SentCollReq}).
			Do(ctx)

		return err
	}

	index := -1
	for i := range usr.ReceivedCollReq {
		if usr.SentCollReq[i] == reqID {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.ReceivedCollReq = append(usr.ReceivedCollReq[:index], usr.ReceivedCollReq[index+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedCollReq": usr.ReceivedCollReq}).
		Do(ctx)

	return err
}

//AppendColleague ... appends to collegue array within user
//takes in eclient, user ID, and collegue ID
func AppendColleague(eclient *elastic.Client, usrID string, colleagueID string) error {

	ctx := context.Background()

	colleagueLock.Lock()
	defer colleagueLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Colleagues = append(usr.Colleagues, colleagueID)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Colleagues": usr.Colleagues}).
		Do(ctx)

	return err

}

//DeleteColleague ... deletes from collegue array within user
//takes in eclient, user ID, and collegue ID
func DeleteColleague(eclient *elastic.Client, usrID string, deleteID string) error {
	ctx := context.Background()

	colleagueLock.Lock()
	defer colleagueLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)
	//idx, err := universal.FindIndex(usr.Colleagues, deleteID) UNIVERSAL PKG
	//temp for-loop:

	index := -1
	for i := range usr.Colleagues {
		if usr.Colleagues[i] == deleteID {
			index = i
		}
	}

	if index < 0 {
		return errors.New("Index non-existent")
	}
	//temp solution stops here

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Colleagues = append(usr.Colleagues[:index], usr.Colleagues[index+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Colleagues": usr.Colleagues}).
		Do(ctx)

	return err

}

func AppendMajorMinor(eclient *elastic.Client, usrID string, majorMinor string, whichOne bool) error {
	//appends to either sent or received collegue request arrays within user
	//takes in eclient, user ID, the major or minor, and a bool
	//true = major, false = minor
	ctx := context.Background()

	procLock.Lock()
	defer procLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.Majors = append(usr.Majors, majorMinor)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"Majors": usr.Majors}).
			Do(ctx)

		return err
	}
	usr.Minors = append(usr.Minors, majorMinor)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	return err

}

//DeleteMajorMinor ... appends to either sent or received collegue request arrays within user
//takes in eclient, user ID, the major or minor, an index of the element within the array, and a bool
//true = major, false = minor
func DeleteMajorMinor(eclient *elastic.Client, usrID string, majorMinor string, whichOne bool) error {

	ctx := context.Background()

	procLock.Lock()
	defer procLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		index := -1
		for i := range usr.Majors {
			if usr.Majors[i] == majorMinor {
				index = i
			}
		}
		if index < 0 {
			return errors.New("Index not found")
		}
		usr.Majors = append(usr.Majors[:index], usr.Majors[index+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"Majors": usr.Majors}).
			Do(ctx)

		return err
	}
	index := -1
	for i := range usr.Minors {
		if usr.Minors[i] == majorMinor {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.Minors = append(usr.Minors[:index], usr.Minors[index+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	return err
}

//AppendFollow ... appends to either sent or received collegue request arrays within user
//takes in eclient, user ID, the follower ID, and a bool
//true = append to following, false = append to followers
func AppendFollow(eclient *elastic.Client, usrID string, followID string, whichOne bool) error {

	ctx := context.Background()

	followLock.Lock()
	defer followLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.Following = append(usr.Following, followID)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"Following": usr.Following}).
			Do(ctx)

		return err
	}
	usr.Followers = append(usr.Followers, followID)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Followers": usr.Followers}).
		Do(ctx)

	return err
}

//DeleteFollow ... whichOne: true = following
//whichOne: false = followers
//followID does nothing
func DeleteFollow(eclient *elastic.Client, usrID string, followID string, whichOne bool) error {

	ctx := context.Background()

	followLock.Lock()
	defer followLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		index := -1
		for i := range usr.Following {
			if usr.Following[i] == followID {
				index = i
			}
		}
		if index < 0 {
			return errors.New("Index not found")
		}
		usr.Following = append(usr.Following[:index], usr.Following[index+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"Following": usr.Following}).
			Do(ctx)

		return err

	}
	index := -1
	for i := range usr.Followers {
		if usr.Followers[i] == followID {
			index = i
		}
	}
	if index < 0 {
		return errors.New("Index not found")
	}
	usr.Followers = append(usr.Followers[:index], usr.Followers[index+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Followers": usr.Followers}).
		Do(ctx)

	return err
}

//AppendProjReq ... appends to either sent or received project request arrays within user
//takes in eclient, user ID, the project ID, and a bool
//true = append to following, false = append to followers
func AppendProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool) error {
	ctx := context.Background()

	projectLock.Lock()
	defer procLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.SentProjReq = append(usr.SentProjReq, projID)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentProjReq": usr.SentProjReq}).
			Do(ctx)

		return err
	}
	usr.ReceivedProjReq = append(usr.ReceivedProjReq, projID)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
		Do(ctx)

	return err
}

//DeleteProjReq ... whichOne: true = sent
//whichOne: false = received
func DeleteProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool) error {
	ctx := context.Background()

	projectLock.Lock()
	defer projectLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		//universal.FindIndex(usr.SentProjReq, projID)
		//temp solution
		index := -1
		for i := range usr.SentProjReq {
			if usr.SentProjReq[i] == projID {
				index = i
				break
			}
		}

		if index < 0 {
			return errors.New("index does not exist")
		}
		//end of temp solution

		usr.SentProjReq = append(usr.SentProjReq[:index], usr.SentProjReq[index+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentProjReq": usr.SentProjReq}).
			Do(ctx)

		return err
	}
	//universal.FindIndex(usr.ReceivedProjReq, projID)
	//temp solution
	index := 0
	for i := range usr.ReceivedProjReq {
		if usr.ReceivedProjReq[i] == projID {
			index = i
			break
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	//end of temp solution
	usr.ReceivedProjReq = append(usr.ReceivedProjReq[:index], usr.ReceivedProjReq[index+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
		Do(ctx)

	return err
}

//AppendLikedEntryID ... appends to either sent or received project request arrays within user
//takes in eclient, user ID, the project ID, and a bool
//true = append to following, false = append to followers
func AppendLikedEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()

	likeLock.Lock()
	defer likeLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.LikedEntryIDs = append(usr.LikedEntryIDs, entryID)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"LikedEntryIDs": usr.LikedEntryIDs}).
		Do(ctx)

	return err

}

//DeleteLikedEntryID ... whichOne: true = following
//whichOne: false = followers
//followID does nothing
func DeleteLikedEntryID(eclient *elastic.Client, usrID string, likerID string) error {
	ctx := context.Background()

	likeLock.Lock()
	defer likeLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.LikedEntryIDs {
		if usr.LikedEntryIDs[i] == likerID {
			index = i
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	usr.LikedEntryIDs = append(usr.LikedEntryIDs[:index], usr.LikedEntryIDs[index+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"LikedEntryIDs": usr.LikedEntryIDs}).
		Do(ctx)

	return err

}

//AppendProject ... appends new project to user
//takes in eclient, user ID, the project ID, and a bool
func AppendProject(eclient *elastic.Client, usrID string, proj types.ProjectInfo) error {
	ctx := context.Background()

	projectLock.Lock()
	defer projectLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Projects = append(usr.Projects, proj)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Projects": usr.Projects}).
		Do(ctx)

	return err

}

//AppendLink ... appends new link to QuickLinks
func AppendLink(eclient *elastic.Client, usrID string, link types.Link) error {
	ctx := context.Background()

	procLock.Lock()
	defer procLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.QuickLinks = append(usr.QuickLinks, link)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"QuickLinks": usr.QuickLinks}).
		Do(ctx)

	return err

}

//DeleteLink ... deletes QuickLink
func DeleteLink(eclient *elastic.Client, usrID string, link types.Link) error {
	ctx := context.Background()

	procLock.Lock()
	defer procLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.QuickLinks {
		if usr.QuickLinks[i] == link {
			index = i
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}

	usr.QuickLinks = append(usr.QuickLinks[:index], usr.QuickLinks[index+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Quicklinks": usr.QuickLinks}).
		Do(ctx)

	return err

}

//AppendTag ... appends a new tag
func AppendTag(eclient *elastic.Client, usrID string, tag string) error {
	ctx := context.Background()

	tagLock.Lock()
	defer tagLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Tags = append(usr.Tags, tag)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Tags": usr.Tags}).
		Do(ctx)

	return err

}

//DeleteTag ... deletes a tag
func DeleteTag(eclient *elastic.Client, usrID string, tag string, idx int) error {
	ctx := context.Background()

	tagLock.Lock()
	defer tagLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	index := -1
	for i := range usr.Tags {
		if usr.Tags[i] == tag {
			index = i
		}
	}
	if index < 0 {
		return errors.New("index does not exist")
	}
	usr.Tags = append(usr.Tags[:idx], usr.Tags[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Tags": usr.Tags}).
		Do(ctx)

	return err

}

//AppendBlock ... appends to the blocked users array
func AppendBlock(eclient *elastic.Client, usrID string, blockID string, whichOne bool) error {
	ctx := context.Background()

	blockLock.Lock()
	defer blockLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.BlockedUsers = append(usr.BlockedUsers, blockID)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
			Do(ctx)

		return err
	} else {
		usr.BlockedUsers = append(usr.BlockedUsers, blockID)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
			Do(ctx)

		return err
	}

}

//DeleteBlock ... unblocks a user by deleting from the blocked array
func DeleteBlock(eclient *elastic.Client, usrID string, blockID string, whichOne bool) error {
	ctx := context.Background()

	blockLock.Lock()
	defer blockLock.Unlock()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		//idx, err := universal.FindIndex(usr.BlockedUsers, blockID)
		//temp solution
		index := 0
		for i := range usr.BlockedUsers {
			if usr.BlockedUsers[i] == blockID {
				index = i
				break
			}
		}
		if index < 0 {
			return errors.New("index does not exist")
		}
		//temp solution end

		usr.BlockedUsers = append(usr.BlockedUsers[:index], usr.BlockedUsers[index+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
			Do(ctx)

		return err
	} else {
		//idx, err := universal.FindIndex(usr.BlockedBy, blockID)
		//temp solution
		index := 0
		for i := range usr.BlockedBy {
			if usr.BlockedBy[i] == blockID {
				index = i
				break
			}
		}
		if index < 0 {
			return errors.New("index does not exist")
		}
		//temp solution end
		usr.BlockedBy = append(usr.BlockedBy[:index], usr.BlockedBy[index+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
			Do(ctx)

		return err

	}

}

//AppendEntryID ... appends a created entry ID to user
func AppendEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()

	entryLock.Lock()
	defer entryLock.Unlock()

	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.EntryIDs = append(usr.EntryIDs, entryID)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": usr.EntryIDs}).
		Do(ctx)

	return err

}

//DeleteEntryID ...deletes entry ID from user array
func DeleteEntryID(eclient *elastic.Client, usrID string, entryID string, idx int) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	usr.EntryIDs = append(usr.EntryIDs[:idx], usr.EntryIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"EntryIDs": usr.EntryIDs}).
		Do(ctx)

	return err

}

func AppendConvoID(eclient *elastic.Client, usrID string, convoID string) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	usr.ConversationIDs = append(usr.ConversationIDs, convoID)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ConversationIDs": usr.ConversationIDs}).
		Do(ctx)

	return err

}

func DeleteConvoID(eclient *elastic.Client, usrID string, convoID string, idx int) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	usr.ConversationIDs = append(usr.ConversationIDs[:idx], usr.ConversationIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ConversationIDs": usr.ConversationIDs}).
		Do(ctx)

	return err

}

func AppendSearch(eclient *elastic.Client, usrID string, newSearch string) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return err
	}

	usr.SearchHist = append(usr.SearchHist, newSearch)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"SearchHist": usr.SearchHist}).
		Do(ctx)

	return err

}
