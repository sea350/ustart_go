package profile

import (
	"fmt"
	"html/template"
	"net/http"

	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//AddWidget ... After widget form submission adds a widget to database
func AddWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
	}
	username := test1.(string)
	r.ParseForm()

	var data []template.HTML
	var classification int

	if r.FormValue("widgetSubmit") == `0` {
		// text
		title := template.HTML(r.FormValue("customHeader"))
		description := template.HTML(r.FormValue("customContent"))
		data = []template.HTML{title, description}
		classification = 0
	}
	if r.FormValue("widgetSubmit") == `1` {
		//gallery
		image := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{image}
		classification = 1
	}
	if r.FormValue("widgetSubmit") == `2` {
		//calendar WIP
		image := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{image}
		classification = 2
	}
	if r.FormValue("widgetSubmit") == `3` {
		//poll WIP
		image := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{image}
		classification = 3
	}
	if r.FormValue("widgetSubmit") == `4` {
		//poll WIP
		image := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{image}
		classification = 4
	}
	if r.FormValue("widgetSubmit") == `5` {
		//soundcloud
		url := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{url}
		classification = 5
	}
	if r.FormValue("widgetSubmit") == `6` {
		//youtube
		url := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{url}
		classification = 6
	}
	if r.FormValue("widgetSubmit") == `7` {
		//codepen
		codepenID := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{codepenID}
		classification = 7
	}
	if r.FormValue("widgetSubmit") == `8` {
		//pinterest
		url := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{url}
		classification = 8
	}
	if r.FormValue("widgetSubmit") == `9` {
		//tumblr
		tumblrEmbedCode := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{tumblrEmbedCode}
		classification = 9
	}
	if r.FormValue("widgetSubmit") == `10` {
		//spoofy
		spotifyEmbedCode := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{spotifyEmbedCode}
		classification = 10
	}
	if r.FormValue("widgetSubmit") == `11` {
		//spoofy
		spotifyEmbedCode := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{spotifyEmbedCode}
		classification = 11
	}
	if r.FormValue("widgetSubmit") == `12` {
		//medium
		username := template.HTML(r.FormValue("UNKNOWN"))
		publication := template.HTML(r.FormValue("UNKNOWN"))
		publicationTag := template.HTML(r.FormValue("UNKNOWN"))
		count := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{username, publication, publicationTag, count}
		classification = 12
	}
	if r.FormValue("widgetSubmit") == `13` {
		//devianart
		username := template.HTML(r.FormValue("UNKNOWN"))
		count := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{username, count}
		classification = 13
	}
	if r.FormValue("widgetSubmit") == `14` {
		//devianart
		username := template.HTML(r.FormValue("UNKNOWN"))
		data = []template.HTML{username}
		classification = 14
	}

	newWidget := types.Widget{UserID: session.Values["DocID"].(string), Data: data, Classification: classification}

	if r.FormValue("editID") == `0` {
		err := uses.AddWidget(eclient, session.Values["DocID"].(string), newWidget)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addWidget.go 128")
		}
	} else {
		err := post.ReindexWidget(eclient, r.FormValue("editID"), newWidget)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addWidget.go 134")
		}
	}

	//contentArray := []rune(comment)
	//username := r.FormValue("username")

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
