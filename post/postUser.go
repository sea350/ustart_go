package post

import (
	"context"
	"errors"

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
                }

                
            }
        }
    }
}`

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

		// Return an error saying ti doesn't exist
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

func AppendToUser(eclient *elastic.Client, usrID string, field string, data interface{}) error {
	return nil
} //RETURN HERE
func RemoveFromUser(eclient *elastic.Client, usrID string, field string, idx int, data interface{}) error {
	return nil
}

//AppendCollReq ...
//  Appends to either sent or received collegue, based on whichOne
//  True = sent; False = received.
func AppendCollReq(eclient *elastic.Client, usrID string, collegueID string, whichOne bool) error {

	ctx := context.Background()
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
func DeleteCollReq(eclient *elastic.Client, usrID string, whichOne bool, idx int) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.SentCollReq = append(usr.SentCollReq[:idx], usr.SentCollReq[idx+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentCollReq": usr.SentCollReq}).
			Do(ctx)

		return err
	}

	usr.ReceivedCollReq = append(usr.ReceivedCollReq[:idx], usr.ReceivedCollReq[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedCollReq": usr.ReceivedCollReq}).
		Do(ctx)

	return err
}

func AppendColleague(eclient *elastic.Client, usrID string, colleagueID string) error {
	//appends to collegue array within user
	//takes in eclient, user ID, and collegue ID
	ctx := context.Background()
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

func DeleteColleague(eclient *elastic.Client, usrID string, idx int) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	usr.Colleagues = append(usr.Colleagues[:idx], usr.Colleagues[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Colleagues": usr.Colleagues}).
		Do(ctx)

	return err

}

func AppendMajorMinor(eclient *elastic.Client, usrID string, major_minor string, whichOne bool) error {
	//appends to either sent or received collegue request arrays within user
	//takes in eclient, user ID, the major or minor, and a bool
	//true = major, false = minor
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)

	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.Majors = append(usr.Majors, major_minor)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"Majors": usr.Majors}).
			Do(ctx)

		return err
	}
	usr.Minors = append(usr.Minors, major_minor)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	return err

}

func DeleteMajorMinor(eclient *elastic.Client, usrID string, major_minor string, whichOne bool, idx int) error {
	//appends to either sent or received collegue request arrays within user
	//takes in eclient, user ID, the major or minor, an index of the element within the array, and a bool
	//true = major, false = minor
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.Majors = append(usr.Majors[:idx], usr.Majors[idx+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"Majors": usr.Majors}).
			Do(ctx)

		return err
	}
	usr.Minors = append(usr.Minors[:idx], usr.Minors[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Minors": usr.Minors}).
		Do(ctx)

	return err
}

func AppendFollow(eclient *elastic.Client, usrID string, followID string, whichOne bool) error {
	//appends to either sent or received collegue request arrays within user
	//takes in eclient, user ID, the follower ID, and a bool
	//true = append to following, false = append to followers
	ctx := context.Background()
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

func DeleteFollow(eclient *elastic.Client, usrID string, whichOne bool, idx int) error {
	//whichOne: true = following
	//whichOne: false = followers
	//followID does nothing
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.Following = append(usr.Following[:idx], usr.Following[idx+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"Following": usr.Following}).
			Do(ctx)

		return err

	}
	usr.Followers = append(usr.Followers[:idx], usr.Followers[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Followers": usr.Followers}).
		Do(ctx)

	return err
}

func AppendProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool) error {
	ctx := context.Background()
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

func DeleteProjReq(eclient *elastic.Client, usrID string, projID string, whichOne bool, idx int) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.SentProjReq = append(usr.SentProjReq[:idx], usr.SentProjReq[idx+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"SentProjReq": usr.SentProjReq}).
			Do(ctx)

		return err
	}
	usr.ReceivedProjReq = append(usr.ReceivedProjReq[:idx], usr.ReceivedProjReq[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"ReceivedProjReq": usr.ReceivedProjReq}).
		Do(ctx)

	return err
}

func AppendLikedEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()
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

func DeleteLikedEntryID(eclient *elastic.Client, usrID string, likerID string) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	idx := 0
	for i := range usr.LikedEntryIDs {
		if usr.LikedEntryIDs[i] == likerID {
			idx = i
		}
	}
	usr.LikedEntryIDs = append(usr.LikedEntryIDs[:idx], usr.LikedEntryIDs[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"LikedEntryIDs": usr.LikedEntryIDs}).
		Do(ctx)

	return err

}

func AppendProject(eclient *elastic.Client, usrID string, proj types.ProjectInfo) error {
	ctx := context.Background()
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

func AppendLink(eclient *elastic.Client, usrID string, link types.Link) error {
	ctx := context.Background()
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

func DeleteLink(eclient *elastic.Client, usrID string, link types.Link, idx int) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	usr.QuickLinks = append(usr.QuickLinks[:idx], usr.QuickLinks[idx+1:]...)

	_, err = eclient.Update().
		Index(esUserIndex).
		Type(esUserType).
		Id(usrID).
		Doc(map[string]interface{}{"Quicklinks": usr.QuickLinks}).
		Do(ctx)

	return err

}

func AppendTag(eclient *elastic.Client, usrID string, tag string) error {
	ctx := context.Background()
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

func DeleteTag(eclient *elastic.Client, usrID string, tag string, idx int) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
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

func AppendBlock(eclient *elastic.Client, usrID string, blockID string, whichOne bool) error {
	ctx := context.Background()
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

func DeleteBlock(eclient *elastic.Client, usrID string, blockID string, idx int, whichOne bool) error {
	ctx := context.Background()
	usr, err := get.GetUserByID(eclient, usrID)
	if err != nil {
		return errors.New("User does not exist")
	}

	if whichOne == true {
		usr.BlockedUsers = append(usr.BlockedUsers[:idx], usr.BlockedUsers[idx+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedUsers": usr.BlockedUsers}).
			Do(ctx)

		return err
	} else {
		usr.BlockedBy = append(usr.BlockedBy[:idx], usr.BlockedBy[idx+1:]...)

		_, err = eclient.Update().
			Index(esUserIndex).
			Type(esUserType).
			Id(usrID).
			Doc(map[string]interface{}{"BlockedBy": usr.BlockedBy}).
			Do(ctx)

		return err

	}

}

func AppendEntryID(eclient *elastic.Client, usrID string, entryID string) error {
	ctx := context.Background()
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
