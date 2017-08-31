package uses

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	post "github.com/sea350/ustart_go/post"
	//get "github.com/sea350/ustart_go/get"
	//"net/http"
	//"io"
	//"fmt"
	//"bytes"
	//"errors"
	//"golang.org/x/crypto/bcrypt"
	"time"
)

func CreateProject (eclient *elastic.Client, title string, description []rune, makerID string)(error){
	var newProj types.Project
	newProj.Name = title
	newProj.Description = description
	newProj.Visible = true
	newProj.CreationDate = time.Now()

	var maker types.Member
	maker.JoinDate = time.Now()
	maker.MemberID = makerID
	maker.Role = 0
	maker.Title = "Creator"
	maker.Visible = true

	newProj.Members = append(newProj.Members, maker)

	err := post.IndexProject(eclient, newProj)
	return err

}

