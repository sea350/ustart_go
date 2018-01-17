package profile

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sea350/ustart_go/middleware/stringHTML"
	types "github.com/sea350/ustart_go/types"
	uses "github.com/sea350/ustart_go/uses"
)

//WallLoad ... Iunno
func WallLoad(w http.ResponseWriter, r *http.Request) {
	// If followingStatus = no
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["DocID"]
	if test1 == nil {
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	fmt.Println(test1)
	fmt.Println("This is debug text, WallLoad.go: 20")
	r.ParseForm()
	entryIDs := r.FormValue("entryIDs")
	fmt.Println(entryIDs)
	var jEntries []types.JournalEntry

	pageID := r.FormValue("pageID")
	if strings.Compare("null", entryIDs) != 0 {

		actualIDs := strings.Split(entryIDs, ",")
		//jEntriesPointer := &jEntries
		jEntries, err5 := uses.LoadEntries(eclient, actualIDs)
		fmt.Println(jEntries[0].ElementID)
		if err5 != nil {
			fmt.Println(err5)
			fmt.Println("This is an error, WallLoad.go: 34")
		}
	}
	var output string
	DocID := session.Values["DocID"].(string)

	output += stringHTML.WallLoadStart(DocID, pageID)

	if strings.Compare("null", entryIDs) != 0 {
		sum := 0
		class0 := `<div class="panel panel-default wallAppend">`

		for i := len(jEntries) - 1; i >= 0; i-- {
			sum++
			if jEntries[i].Element.Classification == 0 {

				class0 += stringHTML.AddClass0Entry(
					jEntries[i].Image,
					jEntries[i].FirstName,
					string(jEntries[i].Element.Content),
					jEntries[i].ElementID,
					string(jEntries[i].NumLikes),
					string(jEntries[i].NumReplies),
					string(jEntries[i].NumShares))
			}
			if jEntries[i].Element.Classification == 2 {
				postIDArray := []string{jEntries[i].Element.ReferenceEntry} // just an array with 1 entry
				jEntry, err5 := uses.LoadEntries(eclient, postIDArray)
				if err5 != nil {
					fmt.Println(err5)
					fmt.Println("This is an error, WallLoad.go: 207")
				}

				class0 += stringHTML.AddClass2Entry(
					string(jEntries[i].Element.Content), //comment
					jEntry[0].Image,
					jEntry[0].FirstName,
					jEntry[0].LastName,
					string(jEntry[0].Element.Content), //bodytext
					string(jEntry[0].ElementID),
					string(jEntry[0].NumLikes),
					string(jEntry[0].NumReplies),
					string(jEntry[0].NumShares),
				)
			}

		}
		class0 += "</div>"
		output += class0
	}

	output += stringHTML.WallLoadEnd(DocID, pageID)

	fmt.Fprintln(w, output)
}
