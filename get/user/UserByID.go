package get

import (
	"context"
	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
	//post "github.com/sea350/ustart_go/post"
)

//UserByID ...
func UserByID(eclient *elastic.Client, userID string) (types.User, error) {
	ctx := context.Background()         //intialize context background
	var usr types.User                  //initialize type user
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.UserIndex).
						Type(globals.UserType).
						Id(userID).
						Do(ctx)

	if err != nil {
		return usr, err
	}

	Err := json.Unmarshal(*searchResult.Source, &usr) //unmarshal type RawMessage into user struct

	return usr, Err

}

/*func GetUserEmailByID(eclient *elastic.Client, usrID string) (string,error) {
	retEmail:=""
	usr, err:= GetUserById(eclient, usrID)

	if (err != nil) {return retEmail,err}
	retEmail = usr.Email

	return retEmail, err
}*/
