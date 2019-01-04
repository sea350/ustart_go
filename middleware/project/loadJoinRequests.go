package project

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/sea350/ustart_go/uses"

	get "github.com/sea350/ustart_go/get/project"
	"github.com/sea350/ustart_go/middleware/client"
	types "github.com/sea350/ustart_go/types"
)

//LoadJoinRequests ...
func LoadJoinRequests(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	ID := r.FormValue("projID") //projectID
	var heads []types.FloatingHead

	proj, err := get.ProjectByID(client.Eclient, ID)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	for index, userID := range proj.MemberReqReceived {
		head, err := uses.ConvertUserToFloatingHead(client.Eclient, userID)
		if err != nil {
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s | index %d", err, index)
		}
		heads = append(heads, head)
	}

	data, err := json.Marshal(heads)
	if err != nil {

		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	fmt.Fprintln(w, string(data))

}
