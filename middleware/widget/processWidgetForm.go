package widget

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/types"
)

//ProcessWidgetForm ... Populates a barebones widget with form data
func ProcessWidgetForm(r *http.Request) (types.Widget, error) {

	var data []template.HTML
	var classification int
	var newWidget types.Widget

	if r.FormValue("widgetSubmit") == `0` {
		// text
		title := template.HTML(r.FormValue("customHeader"))
		description := template.HTML(r.FormValue("customContent"))
		data = []template.HTML{title, description}
		classification = 0
	}
	if r.FormValue("widgetSubmit") == `1` {
		//gallery
		input := template.HTML(r.FormValue("UNKNOWN"))
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
		img := r.FormValue("UNKNOWN")

		image := template.HTML(img)
		data = []template.HTML{image}
		classification = 3
	}
	if r.FormValue("widgetSubmit") == `4` {
		//instagram -- Takes in an instagram post URL
		insta := r.FormValue("instagramInput")
		edit := r.FormValue("editID")

		regX := regexp.MustCompile(`https?:\/\/www\.instagram\.com\/p\/[A-Za-z0-9\-\_]{11}\/.*`)
		if !regX.MatchString(insta) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid URL

		input := template.HTML(insta)
		if edit != `0` {
			widget, err := get.WidgetByID(client.Eclient, edit)
			if err != nil {
				fmt.Println("this is an err, addwidget 66")
				return newWidget, err
			}

			data = append(widget.Data, input)
		} else {
			data = []template.HTML{input}
		}
		classification = 4
	}
	if r.FormValue("widgetSubmit") == `5` {
		//soundcloud -- Takes in a Embed Code
		soundCloud := r.FormValue("scInput")

		regX := regexp.MustCompile(`<iframe width="100%" height="450" scrolling="no" frameborder="no" allow="autoplay" src="https:\/\/w\.soundcloud\.com\/player\/\?url=https%3A\/\/api\.soundcloud\.com\/users\/[0-9]{9}&amp;color=%23[0-9a-f]{6}&amp;auto_play=false&amp;hide_related=false&amp;show_comments=true&amp;show_user=true&amp;show_reposts=false&amp;show_teaser=true"\>\<\/iframe\>`)
		if !regX.MatchString(soundCloud) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		url := template.HTML(soundCloud)
		data = []template.HTML{url}
		classification = 5
	}
	if r.FormValue("widgetSubmit") == `6` {
		//youtube -- Takes in a URL
		yooToob := r.FormValue("ytInput")
		/*
			regX := regexp.MustCompile(``)
			if !regX.MatchString(yooToob) {
				return newWidget, errors.New(`Invalid widget embed code`)
			} //Check valid embed code
		*/
		url := template.HTML(yooToob)
		data = []template.HTML{url}
		classification = 6
	}
	if r.FormValue("widgetSubmit") == `7` {
		//codepen -- Embed code
		codepen := r.FormValue("codepenInput")

		regX := regexp.MustCompile(`<p data-height="265" data-theme-id="0" data-slug-hash="[A-Za-z]{6}" data-default-tab="css,result" data-user="[^"]+" data-embed-version="2" data-pen-title="[^"]+" class="codepen"\>[^<]+\<a href="https:\/\/codepen\.io\/[^\/]+\/pen\/[A-Za-z]{6}\/">[^<]+<\/a> by [^(]+\(<a href="https:\/\/codepen\.io\/[^"]+">@[^<]+<\/a>\) on <a href="https:\/\/codepen\.io">CodePen<\/a>\.<\/p>\n<script async src="https:\/\/static\.codepen\.io\/assets\/embed\/ei\.js"><\/script>`)
		if !regX.MatchString(codepen) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		codepenID := template.HTML(codepen)
		data = []template.HTML{codepenID}
		classification = 7
	}
	if r.FormValue("widgetSubmit") == `8` {
		//pinterest -- Currently will take in user profiles and NOT POSTS!!!!!!!
		pinput := r.FormValue("pinInput")

		regX := regexp.MustCompile(`https:\/\/www\.pinterest\.com\/.+\/`)
		if !regX.MatchString(pinput) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		url := template.HTML(pinput)
		data = []template.HTML{url}

		classification = 8
	}
	if r.FormValue("widgetSubmit") == `9` {
		//tumblr -- Requires a username, no in-depth check needed
		tumblr := r.FormValue("tumblrInput")

		regX := regexp.MustCompile(`[A-Za-z0-9\-\.]{1,32}`)
		if !regX.MatchString(tumblr) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		tumblrEmbedCode := template.HTML(tumblr)
		data = []template.HTML{tumblrEmbedCode}
		classification = 9
	}
	if r.FormValue("widgetSubmit") == `10` {
		//spoofy -- Embed code
		spoofy := r.FormValue("spotInput")
		edit := r.FormValue("editID")

		regX := regexp.MustCompile(`<iframe src="https:\/\/open\.spotify\.com\/embed\/[^"]+" width="300" height="380" frameborder="0" allowtransparency="true"><\/iframe>`)
		if !regX.MatchString(spoofy) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		spotifyEmbedCode := template.HTML(spoofy)
		if edit != `0` {
			widget, err := get.WidgetByID(client.Eclient, edit)
			if err != nil {
				fmt.Println("this is an err, addwidget 156")
				return newWidget, err
			}

			data = append(widget.Data, spotifyEmbedCode)
		} else {
			data = []template.HTML{spotifyEmbedCode}
		}
		classification = 10
	}
	if r.FormValue("widgetSubmit") == `11` {
		//anchor -- Requires link that's almost impossible to get
		ank := r.FormValue("arInput")
		edit := r.FormValue("editID")

		regX := regexp.MustCompile(`https:\/\/anchor\.fm\/[^\/]*\/episodes\/.+`)
		if !regX.MatchString(ank) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		input := template.HTML(ank)
		if r.FormValue("editID") != `0` {
			widget, err := get.WidgetByID(client.Eclient, edit)
			if err != nil {
				fmt.Println("this is an err, addwidget 180")
				return newWidget, err
			}

			data = append(widget.Data, input)
		} else {
			data = []template.HTML{input}
		}
		classification = 11
	}
	if r.FormValue("widgetSubmit") == `12` {
		//medium
		medUsername := r.FormValue("medInput")

		regX := regexp.MustCompile(`[0-9A-Za-z\-]{1,32}`)
		if !regX.MatchString(medUsername) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		username := template.HTML(medUsername)
		publication := template.HTML(r.FormValue("medInput2"))
		publicationTag := template.HTML(r.FormValue("medInput3"))
		count := template.HTML(r.FormValue("medInput4"))
		data = []template.HTML{username, publication, publicationTag, count}
		classification = 12
	}
	if r.FormValue("widgetSubmit") == `13` {
		//devianart -- takes in a username
		da1 := r.FormValue("daInput")
		da2 := r.FormValue("daInput2")

		regX := regexp.MustCompile(`[0-9A-Za-z\-]{1,32}`)
		if !regX.MatchString(da1) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code
		if !regX.MatchString(da2) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		username := template.HTML(da1)
		count := template.HTML(da2)
		data = []template.HTML{username, count}
		classification = 13
	}
	if r.FormValue("widgetSubmit") == `14` {
		//twitch.tv :) -- Takes in a username
		twitch := r.FormValue("twitchInput")

		regX := regexp.MustCompile(`[0-9A-Za-z_]{1,25}`)
		if !regX.MatchString(twitch) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		username := template.HTML(twitch)
		data = []template.HTML{username}
		classification = 14
	}
	if r.FormValue("widgetSubmit") == `15` {
		//calendar widget
		input := r.FormValue("gCalEmbed")

		regX := regexp.MustCompile("(^[a-zA-Z0-9.!#$%&’*+/=?^_`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$)|((<iframe src=\x22)?https:\/\/calendar\.google\.com\/calendar\/((embed)|(htmlembed))\?src=.*)`)
		if !regX.MatchString(input) {
			return newWidget, errors.New(`Invalid widget embed code`)
		} //Check valid embed code

		calendarInput := template.HTML(input)

		data = []template.HTML{calendarInput}
		classification = 15
	}
	if r.FormValue("widgetSubmit") == `16` {
		//github widget username
		urname := r.FormValue("username")

		regX := regexp.MustCompile(`[0-9A-Za-z\-]{1,32}`)
		if !regX.MatchString(urname) {
			return newWidget, errors.New(`Invalid widget embed code`)
		}
		username := template.HTML(urname)
		data = []template.HTML{username}
		classification = 16
	}

	return types.Widget{Data: data, Classification: classification}, nil
}
