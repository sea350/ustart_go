package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
	post "github.com/sea350/ustart_go/post/user"
	"github.com/sea350/ustart_go/types"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {

	ctx := context.Background()

	maq := elastic.NewMatchQuery("Tags", "Project Coaching")
	res, err := eclient.Search().
		Index(globals.UserIndex).
		Type(globals.UserType).
		Query(maq).
		Size(500).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	for _, id := range res.Hits.Hits {
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}

	}

	pc := []string{}
	for _, id := range res.Hits.Hits {
		newSkills := []string{}
		skills := make(map[string]int)
		data := types.User{}
		err = json.Unmarshal(*id.Source, &data)
		if err != nil {
			fmt.Println(err)
		}
		pc = append(pc, data.Email)
		for _, s := range data.Tags {
			skills[s] = 1
		}
		fmt.Println(data.FirstName, data.LastName)
		for k := range skills {
			newSkills = append(newSkills, k)
		}
		fmt.Println(newSkills)
		fmt.Println("---------------------")
		err := post.UpdateUser(eclient, id.Id, "Tags", newSkills)
		if err != nil {
			fmt.Println(err)
		}
	}

}
