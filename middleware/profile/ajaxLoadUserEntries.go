package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	client "github.com/sea350/ustart_go/middleware/client"
	scrollpkg "github.com/sea350/ustart_go/properloading"
)

//AjaxLoadUserEntries ... pulls all entries for a given user and fprints it back as a json array
func AjaxLoadUserEntries(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}

	wallID := r.FormValue("userID")
	if wallID == `` {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println("WARNING: docID not received")
	}
	scrollID := r.FormValue("scrollID")

	res, entries, total, err := scrollpkg.ScrollPageUser(client.Eclient, wallID, scrollID)
	if err != nil {
		if err != io.EOF {
			log.SetFlags(log.LstdFlags | log.Lshortfile)
			log.Println(err) //we might need special treatment for EOF error
		}
	}

	results := make(map[string]interface{})
	results["JournalEntries"] = entries
	results["ScrollID"] = res
	results["TotalHits"] = total

	data, err := json.Marshal(results)
	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}

	fmt.Fprintln(w, string(data))
}
