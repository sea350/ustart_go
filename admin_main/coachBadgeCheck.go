package main

import (

	// admin "github.com/sea350/ustart_go/admin"

	"context"
	"encoding/json"
	"fmt"

	elastic "github.com/olivere/elastic"
	"github.com/sea350/ustart_go/globals"
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

	maq2 := elastic.NewMatchQuery("ID", "UCOACH")
	res2, err := eclient.Search().
		Index(globals.BadgeIndex).
		Type(globals.BadgeType).
		Query(maq2).
		Size(500).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	skills := make(map[string]int{})

	pc := []string{}
	for _, id := range res.Hits.Hits {
		skills := make(map[string]int{})
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
			fmt.Println(k)
		}

	}

	pc2 := []string{}
	for _, id := range res2.Hits.Hits {
		badgeData := types.Badge{}
		err = json.Unmarshal(*id.Source, &badgeData)
		if err != nil {
			fmt.Println(err)
		}
		pc2 = badgeData.Roster
	}
	// ver := int(res.TotalHits())
	// notVer := int(res2.TotalHits())
	// fmt.Println("Total:", ver+notVer)
	// fmt.Println("Project Coaching:", ver)
	// fmt.Println("Coaching:", notVer)

	fmt.Println("Project Coaching:")
	for _, e := range pc {
		fmt.Println(e)
	}
	fmt.Println("--------------------------------")
	fmt.Println("The following used the code: ")
	for _, e := range pc2 {
		fmt.Println(e)
	}

}
