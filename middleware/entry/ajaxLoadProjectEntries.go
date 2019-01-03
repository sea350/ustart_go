package entry

import (
	"encoding/json"
	"fmt"
	"io"
	
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadProjectEntries ... pulls all entries for a given project and fprints it back as a json array
func AjaxLoadProjectEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	docID, _ := session.Values["DocID"]
	if docID == nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	wallID := r.FormValue("userID")
	if wallID == `` {
		
				client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | "+"WARNING: docID not received")
	}
	scrollID := r.FormValue("scrollID")

	res, entries, total, err := scrollpkg.ScrollPageProject(client.Eclient, wallID, docID.(string), scrollID)
	if err != nil {
		if err != io.EOF {
			
			client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
		}
	}

	results := make(map[string]interface{})
	results["JournalEntries"] = entries
	results["ScrollID"] = res
	results["TotalHits"] = total

	data, err := json.Marshal(results)
	if err != nil {
		
		client.Logger.Println("DocID: "+session.Values["DocID"].(string)+" | err: %s", err)
	}

	fmt.Fprintln(w, string(data))
}
