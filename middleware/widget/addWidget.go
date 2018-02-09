package widget

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//AddWidget ... After widget form submission adds a widget to database
func AddWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
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
		input := template.HTML(r.FormValue("instagramInput"))
		data = []template.HTML{input}
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
		//instagram
		input := template.HTML(r.FormValue("instagramInput"))
		if r.FormValue("editID") != `0` {
			widget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
			if err != nil {
				fmt.Println(err)
				fmt.Println("this is an err, addwidget 61")
			}

			data = append(widget.Data, input)
		} else {
			data = []template.HTML{input}
		}
		classification = 4
	}
	if r.FormValue("widgetSubmit") == `5` {
		//soundcloud
		url := template.HTML(r.FormValue("scInput"))
		data = []template.HTML{url}
		classification = 5
	}
	if r.FormValue("widgetSubmit") == `6` {
		//youtube
		url := template.HTML(r.FormValue("ytInput"))
		data = []template.HTML{url}
		classification = 6
	}
	if r.FormValue("widgetSubmit") == `7` {
		//codepen
		codepenID := template.HTML(r.FormValue("codepenInput"))
		data = []template.HTML{codepenID}
		classification = 7
	}
	if r.FormValue("widgetSubmit") == `8` {
		//pinterest
		url := template.HTML(r.FormValue("pinInput"))
		if r.FormValue("editID") != `0` {
			widget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
			if err != nil {
				fmt.Println(err)
				fmt.Println("this is an err, addwidget 96")
			}

			data = append(widget.Data, url)
		} else {
			data = []template.HTML{url}
		}
		classification = 8
	}
	if r.FormValue("widgetSubmit") == `9` {
		//tumblr
		tumblrEmbedCode := template.HTML(r.FormValue("tumblrInput"))
		data = []template.HTML{tumblrEmbedCode}
		classification = 9
	}
	if r.FormValue("widgetSubmit") == `10` {
		//spoofy
		spotifyEmbedCode := template.HTML(r.FormValue("spotInput"))
		if r.FormValue("editID") != `0` {
			widget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
			if err != nil {
				fmt.Println(err)
				fmt.Println("this is an err, addwidget 108")
			}

			data = append(widget.Data, spotifyEmbedCode)
		} else {
			data = []template.HTML{spotifyEmbedCode}
		}
		classification = 10
	}
	if r.FormValue("widgetSubmit") == `11` {
		//anchor
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
		//twitch.tv :)
		username := template.HTML(r.FormValue("twitchInput"))
		data = []template.HTML{username}
		classification = 14
	}
	if r.FormValue("widgetSubmit") == `15` {
		//skills
		//special class
		tags := strings.Split(",", r.FormValue("tagInput"))
		if r.FormValue("editID") != `0` {
			widget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
			if err != nil {
				fmt.Println(err)
				fmt.Println("this is an error: middleware/profile/addWidget.go 151")
			}
			data = widget.Data
			for _, tag := range tags {
				data = append(data, template.HTML(tag))
			}
		} else {
			for _, tag := range tags {
				data = append(data, template.HTML(tag))
			}
		}
		classification = 15
	}

	newWidget := types.Widget{UserID: session.Values["DocID"].(string), Data: data, Classification: classification}

	if r.FormValue("editID") == `0` {
		err := uses.AddWidget(client.Eclient, session.Values["DocID"].(string), newWidget)
		fmt.Println("Widget added. Classification:")
		fmt.Println(newWidget.Classification)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addWidget.go 169")
		}
	} else {
		err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addWidget.go 175")
		}
	}

	//contentArray := []rune(comment)
	//username := r.FormValue("username")

	http.Redirect(w, r, "/profile/"+username, http.StatusFound)
}
