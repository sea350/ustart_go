package widget

import (
	"fmt"
	"html/template"
	"net/http"

	getProj "github.com/sea350/ustart_go/get/project"
	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	post "github.com/sea350/ustart_go/post/widget"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//AddProjectWidget ... After widget form submission adds a widget to database
func AddProjectWidget(w http.ResponseWriter, r *http.Request) {
	session, _ := client.Store.Get(r, "session_please")
	test1, _ := session.Values["Username"]
	if test1 == nil {
		// No username in session
		http.Redirect(w, r, "/~", http.StatusFound)
		return
	}
	checkerEnable := false
	project, err := getProj.ProjectByID(client.Eclient, r.FormValue("projectWidget"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("this is an error: middleware/profile/addProjectWidget.go 2252")
	}

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
		//instagram -- Takes in an instagram post URL
		if checkerEnable {
			checker := uses.StringChecker(r.FormValue("instagramInput"), "instagram.com") //Check valid URL
			if !checker {
				http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
				fmt.Println("invalid widget embed code")
				return
			}
		}
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
		//soundcloud -- Takes in a Soundcloud Embed Code
		if checkerEnable {
			checker := uses.StringChecker(r.FormValue("scInput"), "soundcloud.com") //Check valid Embed
			if !checker {
				http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
				fmt.Println("invalid widget embed code")
				return
			}
		}
		url := template.HTML(r.FormValue("scInput"))
		data = []template.HTML{url}
		classification = 5
	}
	if r.FormValue("widgetSubmit") == `6` {
		//youtube -- Takes in a YouTube URL
		if checkerEnable {
			checker1 := uses.StringChecker(r.FormValue("ytinput"), "youtube.com") //Check valid URL
			if !checker1 {
				http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
				fmt.Println("invalid widget embed code")
				return
			}
		}
		url := template.HTML(r.FormValue("ytInput"))
		data = []template.HTML{url}
		classification = 6
	}
	if r.FormValue("widgetSubmit") == `7` {
		//codepen -- Takes in an Embed code
		if checkerEnable {
			checker := uses.StringChecker(r.FormValue("codepenInput"), "codepen.io") //Check valid Embed
			if !checker {
				http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
				fmt.Println("invalid widget embed code")
				return
			}
		}
		codepenID := template.HTML(r.FormValue("codepenInput"))
		data = []template.HTML{codepenID}
		classification = 7
	}
	if r.FormValue("widgetSubmit") == `8` {
		//pinterest -- Takes in a pinterest USER PROFILE!
		if checkerEnable {
			checker := uses.StringChecker(r.FormValue("pinInput"), "pinterest.com")
			if !checker {
				http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
				fmt.Println("invalid widget embed code")
				return
			}
		}
		url := template.HTML(r.FormValue("pinInput"))
		data = []template.HTML{url}

		classification = 8
	}
	if r.FormValue("widgetSubmit") == `9` {
		//tumblr -- Username
		/*
				if checkerEnable {
				checker := uses.StringChecker(r.FormValue("tumblrInput"), "tumblr.com") //Check valid Embed
				if !checker {
					http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
					fmt.Println("invalid widget embed code")
					return
				}
			}
		*/

		tumblrEmbedCode := template.HTML(r.FormValue("tumblrInput"))
		data = []template.HTML{tumblrEmbedCode}
		classification = 9
	}
	if r.FormValue("widgetSubmit") == `10` {
		//spoofy -- Embed code
		if checkerEnable {
			checker := uses.StringChecker(r.FormValue("spotInput"), "spotify.com") //Check valid Embed
			if !checker {
				http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
				fmt.Println("invalid widget embed code")
				return
			}
		}
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
		//anchor -- Takes in a .fm URL
		if checkerEnable {
			checker := uses.StringChecker(r.FormValue("arInpus"), "anchor.fm") //Check valid Embed
			if !checker {
				http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
				fmt.Println("invalid widget embed code")
				return
			}
		}
		input := template.HTML(r.FormValue("arInput"))
		if r.FormValue("editID") != `0` {
			widget, err := get.WidgetByID(client.Eclient, r.FormValue("editID"))
			if err != nil {
				fmt.Println(err)
				fmt.Println("this is an err, addwidget 134")
			}

			data = append(widget.Data, input)
		} else {
			data = []template.HTML{input}
		}
		classification = 11
	}
	if r.FormValue("widgetSubmit") == `12` {
		//medium
		username := template.HTML(r.FormValue("medInput"))
		publication := template.HTML(r.FormValue("medInput2"))
		publicationTag := template.HTML(r.FormValue("medInput3"))
		count := template.HTML(r.FormValue("medInput4"))
		data = []template.HTML{username, publication, publicationTag, count}
		classification = 12
	}
	if r.FormValue("widgetSubmit") == `13` {
		//devianart -- Takes in a username
		//if checkerEnable {
		// checker := uses.StringChecker(r.FormValue("daInput"), "deviantart.com") //Check valid Embed
		// if !checker {
		// 	http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
		// 	fmt.Println("invalid widget embed code")
		// 	return
		// }
		//}
		username := template.HTML(r.FormValue("daInput"))
		count := template.HTML(r.FormValue("daInput2"))
		data = []template.HTML{username, count}
		classification = 13
	}
	if r.FormValue("widgetSubmit") == `14` {
		//twitch.tv :) -- Takes in a Username
		/*
			if checkerEnable {
				checker := uses.StringChecker(r.FormValue("twitchInput"), "twitch.tv") //Check valid Embed
				if !checker {
					http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
					fmt.Println("invalid widget embed code")
					return
				}
			}
		*/

		username := template.HTML(r.FormValue("twitchInput"))
		data = []template.HTML{username}
		classification = 14
	}

	newWidget := types.Widget{UserID: r.FormValue("projectWidget"), Data: data, Classification: classification}

	if r.FormValue("editID") == `0` {
		fmt.Println("this is debug text middeware/widget/addprojectidget.go")
		fmt.Println(r.FormValue("projectWidget"))
		err := uses.AddWidget(client.Eclient, r.FormValue("projectWidget"), newWidget, true)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addProjectWidget.go 206")
		}
	} else {
		err := post.ReindexWidget(client.Eclient, r.FormValue("editID"), newWidget)
		if err != nil {
			fmt.Println(err)
			fmt.Println("this is an error: middleware/profile/addWidget.go 212")
		}
	}

	http.Redirect(w, r, "/Projects/"+project.URLName, http.StatusFound)
	return
}
