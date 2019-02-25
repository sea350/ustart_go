package main

import (
	// getUser "github.com/sea350/ustart_go/get/user"
	// post "github.com/sea350/ustart_go/post/user"
	// "context"

	"github.com/sea350/ustart_go/types"

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

	// ctx := context.Background()

	var newPrivs = []types.Privilege{SetMemberPrivileges(0), SetMemberPrivileges(1), SetMemberPrivileges(2)}

	fmt.Println(newPrivs)
	// maq := elastic.NewMatchAllQuery()

	// res, err := eclient.Search().
	// 	Index(globals.ProjectIndex).
	// 	Type(globals.ProjectType).
	// 	Query(maq).
	// 	Do(ctx)

	// if err != nil {
	// 	fmt.Println("LINE 40,", err)
	// }

	// for _, hit := res.Hits.Hits{
	// 	err = post.UpdateProject(eclient, hit.Id, "PrivilegeProfiles", newPrivs)
	// 	if err != nil {
	// 		fmt.Println("LINE 42,", err)
	// 	}
	// }

}
