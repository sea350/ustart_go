package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sea350/ustart_go/uses"

	client "github.com/sea350/ustart_go/middleware/client"
)

//AjaxLoadComments ... pulls all entries for a given user and fprints it back as a json array
func AjaxLoadComments(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	parentID := r.FormValue("postID")
	parent, entries, err := uses.LoadComments(client.Eclient, parentID, 0, -1)
	if err != nil {
		fmt.Println("err middleware/profile/ajaxloadcomments line 25")
		fmt.Println(err)
	}

	data, err := json.Marshal(entries)
	if err != nil {
		fmt.Println("err middleware/profile/ajaxloadcomments line 32")
		fmt.Println(err)
	}

	data2, err := json.Marshal(parent)
	if err != nil {
		fmt.Println("err middleware/profile/ajaxloadcomments line 32")
		fmt.Println(err)
	}

	fmt.Fprintln(w, string(data))
	fmt.Fprintln(w, string(data2))
}
