package main

import (
	"fmt"

	// admin "github.com/sea350/ustart_go/admin"

	"github.com/sea350/ustart_go/globals"
	post "github.com/sea350/ustart_go/post/badge"
	"github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

var eclient, _ = elastic.NewSimpleClient(elastic.SetURL(globals.ClientURL))

func main() {
	// err := admin.ModifyBadge(eclient, "USTART", "give", "gl1144@nyu.edu", "")
	// badge, err := get.BadgeByType(eclient, "USTART")

	// err = post.UpdateBadge(eclient, badge.ID, "Roster", append(badge.Roster, "gl1144@nyu.edu"))
	// fmt.Println(err)
	var sdB types.Badge
	sdB.ID = "SDSPRING19B"
	sdB.Type = "Senior Design Spring 2019"
	sdB.ImageLink = "https://s3.amazonaws.com/ustart-default/Student_badge.png"
	sdB.Roster = []string{"jia247@nyu.edu", "sb5290@nyu.edu", "mb6004@nyu.edu",
		"mc5870@nyu.edu", "hui.chiang@nyu.edu", "chourdiaanjali@nyu.edu",
		"dd2390@nyu.edu", "markguindi@nyu.edu", "rga267@nyu.edu", "zh754@nyu.edu",
		"hkh242@nyu.edu", "bh1531@nyu.edu", "csk387@nyu.edu", "jk5149@nyu.edu",
		"jpk389@nyu.edu", "cjlee@nyu.edu", "mel526@nyu.edu", "csl459@nyu.edu",
		"rm4078@nyu.edu", "tmm500@nyu.edu", "hm1487@nyu.edu", "nick.nguyen@nyu.edu",
		"yq544@nyu.edu", "jsr483@nyu.edu", "gr1188@nyu.edu", "iar252@nyu.edu",
		"doron.rasis@nyu.edu", "jeremyrivera@nyu.edu", "vr714@nyu.edu", "uss204@nyu.edu",
		"kps325@nyu.edu", "js8327@nyu.edu", "sps450@nyu.edu", "vsykoralovaas@nyu.edu",
		"hoyin.wan@nyu.edu", "vz365@nyu.edu"}

	sdB.Tags = []string{"CS4513B_S19"}

	////////////////////////////////////////////////////////////////////////
	var sdA types.Badge
	sdA.ID = "SDSPRING19A"
	sdA.Type = "Senior Design Spring 2019"
	sdA.ImageLink = "https://s3.amazonaws.com/ustart-default/Student_badge.png"
	sdA.Roster = []string{"ad3230@nyu.edu", "dcd310@nyu.edu", "ae1389@nyu.edu", "graeme.ferguson@nyu.edu",
		"sepehr.yazdani@nyu.edu", "jwg327@nyu.edu", "ag5278@nyu.edu", "bh1642@nyu.edu",
		"si751@nyu.edu", "ssi256@nyu.edu", "dk3094@nyu.edu", "gk1307@nyu.edu",
		"dk2901@nyu.edu", "ll3087@nyu.edu", "al4596@nyu.edu", "ksl397@nyu.edu",
		"jjl656@nyu.edu", "ll3056@nyu.edu", "mm8088@nyu.edu", "maliat.manzur@nyu.edu",
		"sm6942@nyu.edu", "am7100@nyu.edu", "rdm420@nyu.edu", "mpn272@nyu.edu",
		"mp3685@nyu.edu", "duc.pham@nyu.edu", "rtr266@nyu.edu", "ps3042@nyu.edu",
		"ns2729@nyu.edu", "rs5666@nyu.edu", "vhs238@nyu.edu", "shs474@nyu.edu",
		"anthony.taldone@nyu.edu", "jt2908@nyu.edu", "amv445@nyu.edu", "vo.richardjohn@nyu.edu",
		"ox207@nyu.edu", "ry821@nyu.edu", "ry745@nyu.edu", "ty787@nyu.edu", "chz224@nyu.edu",
		"david.zheng@nyu.edu", "xz1438@nyu.edu"}

	sdA.Tags = []string{"CS4513A_S19"}
	////////////////////////////////////////////////////////////////////////////
	var sweA types.Badge
	sweA.ID = "SWESPRING19A"
	sweA.Type = "Software Engineering Spring 2019"
	sweA.ImageLink = "https://s3.amazonaws.com/ustart-default/Student_badge.png"
	sweA.Roster = []string{"zehra@nyu.edu", "aca432@nyu.edu", "sb5840@nyu.edu", "ab6858@nyu.edu",
		"byc241@nyu.edu", "xc1008@nyu.edu", "plc300@nyu.edu", "hd892@nyu.edu",
		"md3756@nyu.edu", "eg2584@nyu.edu", "karan.ganta@nyu.edu", "cg2738@nyu.edu",
		"ll3325@nyu.edu", "cl3616@nyu.edu", "dl3474@nyu.edu", "al4991@nyu.edu",
		"ln961@nyu.edu", "felicity.ng@nyu.edu", "mr4739@nyu.edu", "asimsatti@nyu.edu",
		"carrollshen@nyu.edu", "ss10198@nyu.edu", "js10022@nyu.edu", "as9365@nyu.edu",
		"blw322@nyu.edu", "sx563@nyu.edu", "jz2456@nyu.edu", "zz1241@nyu.edu"}
	sweA.Tags = []string{"CS4523A_S19"}

	/////////////////////////////////////////////////////////////////////////
	var sweB types.Badge
	sweB.ID = "SWESPRING19B"
	sweB.Type = "Software Engineering Spring 2019"
	sweB.ImageLink = "https://s3.amazonaws.com/ustart-default/Student_badge.png"
	sweB.Roster = []string{"gc1364@nyu.edu", "hansen.chen@nyu.edu", "sc6094@nyu.edu", "kc3247@nyu.edu",
		"md3837@nyu.edu", "ae1495@nyu.edu", "ph1335@nyu.edu", "ai974@nyu.edu",
		"mj1803@nyu.edu", "teddykim@nyu.edu", "jk5541@nyu.edu", "dql202@nyu.edu",
		"ajl607@nyu.edu", "bkl263@nyu.edu", "fnm225@nyu.edu", "kn1173@nyu.edu",
		"an2343@nyu.edu", "dp2387@nyu.edu", "jq524@nyu.edu", "hq343@nyu.edu",
		"sr4742@nyu.edu", "jinzhaosu@nyu.edu", "jgw298@nyu.edu", "samantha.waln@nyu.edu",
		"mw3482@nyu.edu", "jwon@nyu.edu", "bx357@nyu.edu", "kennyyip@nyu.edu",
		"zhangyu@nyu.edu", "nz710@nyu.edu"}
	sweB.Tags = []string{"CS4523B_S19"}

	sdAPrint, err1 := post.IndexBadge(eclient, sdA)
	fmt.Println(sdAPrint, err1)

	sdBPrint, err2 := post.IndexBadge(eclient, sdB)
	fmt.Println(sdBPrint, err2)

	sweAPrint, err3 := post.IndexBadge(eclient, sweA)
	fmt.Println(sweAPrint, err3)

	sweBPrint, err4 := post.IndexBadge(eclient, sweB)
	fmt.Println(sweBPrint, err4)
}
