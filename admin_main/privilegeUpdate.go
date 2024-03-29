package main

import (
	// getUser "github.com/sea350/ustart_go/get/user"
	// post "github.com/sea350/ustart_go/post/user"
	"context"

	"github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"

	post "github.com/sea350/ustart_go/post/project"

	// post "github.com/sea350/ustart_go/post/user"

	// admin "github.com/sea350/ustart_go/admin"

	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	ctx := context.Background()

	var newPrivs = []types.Privileges{uses.SetMemberPrivileges(0), uses.SetMemberPrivileges(1), uses.SetMemberPrivileges(2)}

	fmt.Println(newPrivs)
	maq := elastic.NewMatchAllQuery()

	res, err := eclient.Search().
		Index(globals.ProjectIndex).
		Type(globals.ProjectType).
		Query(maq).
		Size(200).
		Do(ctx)

	if err != nil {
		fmt.Println("LINE 40,", err)
	}

	for _, hit := range res.Hits.Hits {
		err = post.UpdateProject(eclient, hit.Id, "PrivilegeProfiles", newPrivs)
		fmt.Println(hit.Id)
		if err != nil {
			fmt.Println("LINE 42,", err)
		}
	}

}
