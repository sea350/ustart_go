package main

import (
	getUser "github.com/sea350/ustart_go/get/user"

	// getProj "github.com/sea350/ustart_go/get/project"
	// post "github.com/sea350/ustart_go/post/project"

	// post "github.com/sea350/ustart_go/post/user"

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	usr, err := getUser.UserByUsername(eclient, "www.youtu.be/y5gnlptxgb4")
	// usrID, err := getUser.IDByUsername(eclient, "www.youtu.be/y5gnlptxgb4")
	// usr2, err := getUser.UserByUsername(eclient, "min")
	// usrID2, err := getUser.IDByUsername(eclient, "min")

	// fmt.Println(len(usr.QuickLinks))
	// var emp []types.Link
	// err = post.UpdateUser(eclient, usrID, "QuickLinks", emp)
	if err != nil {
		fmt.Println("LINE 24,", err)
	}

	// err = post.UpdateUser(eclient, usrID2, "QuickLinks", emp)
	// if err != nil {
	// 	fmt.Println("LINE 24,", err)
	// }

	fmt.Println(usr.FirstName, usr.LastName)
}

// func main() {

// 	proj, err := getProj.ProjectByURL(eclient, "ITSpackages")
// 	projID, err := getProj.ProjectIDByURL(eclient, "ITSpackages")
// 	fmt.Println(len(proj.QuickLinks))
// 	var emp []types.Link
// 	err = post.UpdateProject(eclient, projID, "QuickLinks", emp)
// 	if err != nil {
// 		fmt.Println("LINE 51,", err)
// 	}

// }
