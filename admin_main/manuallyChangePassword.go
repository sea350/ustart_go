package main

import (
	"fmt"

	elastic "github.com/olivere/elastic"
	getUser "github.com/sea350/ustart_go/get/user"
	globals "github.com/sea350/ustart_go/globals"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
	"golang.org/x/crypto/bcrypt"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

// newProj.Avatar = "https://i.imgur.com/8BnkFLO.png"
// newProj.Banner = "https://i.imgur.com/XTj1t1J.png"
// newUsr.Avatar = "https://i.imgur.com/8BnFLO.png"
// newUsr.Banner = "https://i.imgur.com/XTj1t1J.png"

// //////////////////////////////////////////////
// newProj.Avatar = "https://ustart-default.s3.amazonaws.com/Defult_Project_Page_Logo.png"
// newProj.Banner = "https://ustart-default.s3.amazonaws.com/Defult_Profile_Banner_Logo.png"
// newUsr.Avatar = "https://ustart-default.s3.amazonaws.com/Defult_Profile_Page_Logo.png"
// newUsr.Banner = "https://ustart-default.s3.amazonaws.com/Defult_Profile_Banner_Logo.png"

func main() {

	usr, err := getUser.UserByEmail(eclient, "ronald11341@aol.com")

	fmt.Println("Get user by email error", err)

	userID, err := getUser.UseIDByUsername(eclient, usr.Username)

	fmt.Println("Get userID by username error", err)
	newPass := []byte("adanat11341")
	newHashedPass, err := bcrypt.GenerateFromPassword(newPass, bcrypt.DefaultCost)
	// err = post.UpdateUser(eclient, userID, "Password", newHashedPass)
	usr.Password = newHashedPass

	clearWarnings := make(map[string]types.LoginWarning)
	usr.LoginWarnings = clearWarnings
	//Clear login lockout counter
	if err == nil {
		err = post.ReindexUser(eclient, userID, usr)
	}

	fmt.Println("Reindex user error", err)

}
