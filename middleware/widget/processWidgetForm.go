package widget

import (
	"encoding/json"
	"errors"
	"html/template"

	"net/http"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	get "github.com/sea350/ustart_go/get/widget"
	client "github.com/sea350/ustart_go/middleware/client"
	"github.com/sea350/ustart_go/types"
	"github.com/sea350/ustart_go/uses"
)

//ProcessWidgetForm ... Populates a barebones widget with form data
func ProcessWidgetForm(r *http.Request) (types.Widget, error) {

	checkerEnable := false

	var data []template.HTML
	var classification int
	var newWidget types.Widget

	edit := r.FormValue("editID")

	if r.FormValue("widgetSubmit") == `0` {
		// text
		p := bluemonday.UGCPolicy()
		html := p.Sanitize(r.FormValue("customHeader"))
		title := template.HTML(html)
		htmlDesc := p.Sanitize(r.FormValue("customContent"))
		description := template.HTML(htmlDesc)
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
		img := r.FormValue("UNKNOWN")

		image := template.HTML(img)
		data = []template.HTML{image}
		classification = 3
	}

	if r.FormValue("widgetSubmit") == `4` {
		//instagram -- Takes in an instagram post URL

		insta := r.FormValue("instagramInput")

		// regX := regexp.MustCompile(`https?:\/\/www\.instagram\.com\/p\/[A-Za-z0-9\-\_]{11}\/*`)

		// if !regX.MatchString(insta) {
		// 	return newWidget, errors.New(`Unusable Instagram URL`)
		// } //Check valid URL

		testArray := []string{}
		err := json.Unmarshal([]byte(insta), &testArray)
		if err != nil {
			input := template.HTML(insta)
			if edit != `0` {
				widget, err := get.WidgetByID(client.Eclient, edit)
				if err != nil {

					client.Logger.Println("err: ", err)
					return newWidget, err
				}

				data = append(widget.Data, input)
			} else {
				data = []template.HTML{input}
			}
		} else {
			for _, elem := range testArray {
				data = append(data, template.HTML(elem))
			}
		}

		classification = 4

	}
	if r.FormValue("widgetSubmit") == `5` {
		//soundcloud -- Takes in a Embed Code
		soundCloud := r.FormValue("scInput")

		regX := regexp.MustCompile(`<iframe width="[0-9%]{0,4}" height="[0-9%]{0,4}" scrolling="[a-z]+" frameborder="[a-z]+" allow="autoplay" src="https:\/\/w\.soundcloud\.com\/player\/\?url=https%3A\/\/api\.soundcloud\.com\/[a-z]+\/[0-9]{6,9}&color=%23[0-9a-f]{6}&auto_play=[a-z]+&hide_related=[a-z]+&show_comments=[a-z]+&show_user=[a-z]+&show_reposts=[a-z]+&show_teaser=[a-z]+(&visual=[a-z]+)*"><\/iframe>`)
		if !regX.MatchString(soundCloud) {
			return newWidget, errors.New(`Unusable Soundcloud Embed Code`)
		} //Check valid embed code

		url := template.HTML(soundCloud)
		data = []template.HTML{url}
		classification = 5
	}
	if r.FormValue("widgetSubmit") == `6` {
		//youtube -- Takes in a URL
		yooToob := r.FormValue("ytinput")
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
		regX := regexp.MustCompile(`https:\/\/codepen\.io\/[^\/]*\/pen\/.+`)
		if !regX.MatchString(r.FormValue("codepenInput")) {
			return newWidget, errors.New(`Unusable CodePen Embed`)
		}
		if checkerEnable {
			checker := uses.StringChecker(r.FormValue("codepenInput"), "codepen.io") //Check valid Embed
			if !checker {
				return newWidget, errors.New(`Unusable Codepen Embed Code`)
			}
		}
		codepenID := template.HTML(r.FormValue("codepenInput"))
		data = []template.HTML{codepenID}
		classification = 7
	}
	if r.FormValue("widgetSubmit") == `8` {
		//pinterest -- Currently will take in user profiles and NOT POSTS!!!!!!!
		pinput := r.FormValue("pinInput")

		regX := regexp.MustCompile(`https:\/\/www\.pinterest\.com\/.+`)
		if !regX.MatchString(pinput) {
			return newWidget, errors.New(`Unusable Pinterest URL`)
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
			return newWidget, errors.New(`Unusable Tumblr User`)
		} //Check valid embed code

		tumblrEmbedCode := template.HTML(tumblr)
		data = []template.HTML{tumblrEmbedCode}
		classification = 9
	}
	if r.FormValue("widgetSubmit") == `10` {
		//spoofy -- Embed code
		spoofy := r.FormValue("spotInput")

		/*
			regX := regexp.MustCompile(`<iframe src="https:\/\/open\.spotify\.com\/embed\/[^"]+" width="300" height="380" frameborder="0" allowtransparency="true"><\/iframe>`)
			if !regX.MatchString(spoofy) {
				return newWidget, errors.New(`Unusable Spotify Embed Code`)
			} //Check valid embed code
		*/
		testArray := []string{}
		err := json.Unmarshal([]byte(spoofy), &testArray)
		if err != nil {
			input := template.HTML(spoofy)
			if edit != `0` {
				widget, err := get.WidgetByID(client.Eclient, edit)
				if err != nil {

					client.Logger.Println("err: ", err)
					return newWidget, err
				}

				data = append(widget.Data, input)
			} else {
				data = []template.HTML{input}
			}
		} else {
			for _, elem := range testArray {
				data = append(data, template.HTML(elem))
			}
		}
		classification = 10
	}
	if r.FormValue("widgetSubmit") == `11` {
		//anchor -- Requires link that's almost impossible to get
		ank := r.FormValue("arInput")

		/*
			regX := regexp.MustCompile(``)
			if !regX.MatchString(ank) {
				return newWidget, errors.New(`Invalid widget embed code`)
			} //Check valid embed code
		*/
		testArray := []string{}
		err := json.Unmarshal([]byte(ank), &testArray)
		if err != nil {
			input := template.HTML(ank)
			if edit != `0` {
				widget, err := get.WidgetByID(client.Eclient, edit)
				if err != nil {

					client.Logger.Println("err: ", err)
					return newWidget, err
				}

				data = append(widget.Data, input)
			} else {
				data = []template.HTML{input}
			}
		} else {
			for _, elem := range testArray {
				data = append(data, template.HTML(elem))
			}
		}
		classification = 11
	}
	if r.FormValue("widgetSubmit") == `12` {
		//medium
		medUsername := r.FormValue("medInput")

		regX := regexp.MustCompile(`[0-9A-Za-z\-]{1,32}`)
		if !regX.MatchString(medUsername) {
			return newWidget, errors.New(`Unusable Medium Username`)
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

		/*
			if checkerEnable {
				checker := uses.StringChecker(r.FormValue("daInput"), "deviantart.com") //Check valid Embed

				if !checker {
					return newWidget, errors.New(`Invalid widget embed code`)
				}
			}
		*/

		username := template.HTML(r.FormValue("daInput"))
		count := template.HTML(r.FormValue("daInput2"))
		data = []template.HTML{username, count}
		classification = 13
	}
	if r.FormValue("widgetSubmit") == `14` {
		//twitch.tv :) -- Takes in a username
		twitch := r.FormValue("twitchInput")

		regX := regexp.MustCompile(`[0-9A-Za-z_]{1,25}`)
		if !regX.MatchString(twitch) {
			return newWidget, errors.New(`Unusable Twitch Username`)
		} //Check valid embed code

		username := template.HTML(twitch)
		data = []template.HTML{username}
		classification = 14
	}
	if r.FormValue("widgetSubmit") == `15` {
		//calendar widget

		calendarInput := template.HTML(r.FormValue("gCalEmbed"))
		regX := regexp.MustCompile(`([a-zA-Z0-9]+)([.{1}])?([a-zA-Z0-9]+)@gmail([.])com`)
		if !regX.MatchString(calendarInput) {
			return newWidget, errors.New(`Email did not match valid email criteria`)
		} //Check valid embed code
		data = []template.HTML{calendarInput}
		classification = 15
	}
	if r.FormValue("widgetSubmit") == `16` {
		//github widget username

		username := template.HTML(r.FormValue("username"))
		count := template.HTML(r.FormValue("git-count"))
		data = []template.HTML{username, count}
		classification = 16
	}
	if r.FormValue("widgetSubmit") == `17` {
		//gallery widget
		galleryFile, galleryHeader, _ := r.FormFile("galleryImageInput")
		buffer := make([]byte, 512)
		_, _ = galleryFile.Read(buffer)
		defer galleryFile.Close()
		if http.DetectContentType(buffer)[0:5] == "image" || galleryHeader.Size == 0 {
			name := strings.Split(galleryHeader.Filename, ".")
			contents := buffer
			data = []template.HTML{template.HTML(name[0]), template.HTML(contents)}
		}
		classification = 17
	}

	return types.Widget{Data: data, Classification: classification}, nil
}
