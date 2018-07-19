package uses

import (
	"log"
	"os"

	getDash "github.com/sea350/ustart_go/get/dashboard"
	getUser "github.com/sea350/ustart_go/get/user"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//UNFINISHED

//AggregateDashData ...
func AggregateDashData(eclient *elastic.Client, viewerID string) (types.DashboardAggregate, error) {
	var dashData types.DashboardAggregate
	dashData.RequestAllowed = true

	user, err := getUser.UserByID(eclient, viewerID)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}

	data, err := getDash.DashboardByUsername(eclient, user.Username)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		dir, _ := os.Getwd()
		log.Println(dir, err)
	}
	dashData.DashboardData = data

	//Should the dashboard struct have its own DocID?

	return dashData, err

}
