package uses

import(
	elastic "gopkg.in/olivere/elastic.v5"
	//types "github.com/sea350/ustart_go/types"
	post "github.com/sea350/ustart_go/post"
	get "github.com/sea350/ustart_go/get"
	//"fmt"
	"errors"
	//"time"
)

func NukeUser(eclient *elastic.Client, userID string, areYouSure bool) (bool, []error){

	//designed to continue functionality despite encountering errors uncless core gets are intereupted
	nukeSucess := false
	var errorStack []error
	if (!areYouSure){
		return nukeSucess, append(errorStack, errors.New("You are not sure, make a decision and come back"))
	}

	usr, err := get.GetUserByID(eclient, userID)
	if err != nil {
		errorStack = append(errorStack, err)
		errorStack = append(errorStack, errors.New("Problem getting user, no changes were made"))
		return nukeSucess, errorStack
	}

	err = post.UpdateUser(eclient, userID, "Visible", false)
	if err != nil {
		errorStack = append(errorStack, err)
		errorStack = append(errorStack, errors.New("Problem modifying user, no changes were made"))
		return nukeSucess, errorStack
	}

	for _, i := range usr.EntryIDs {
		//sets all entries to invisible
		err := post.UpdateUser(eclient, i, "Visible", false)
		if err != nil {
			errorStack = append(errorStack, err)
			errorStack = append(errorStack, errors.New("Problem modifying entries, changes incomplete"))
		}
	}

	//likes needs no work

	//delete chat here

	//manage projects here

	//followers following?

	//colleagues
	return nukeSucess, errorStack

}