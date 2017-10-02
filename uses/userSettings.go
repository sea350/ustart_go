package uses

import(
	elastic "gopkg.in/olivere/elastic.v5"
	types "github.com/sea350/ustart_go/types"
	post "github.com/sea350/ustart_go/post"
	get "github.com/sea350/ustart_go/get"
	"errors"
	//"golang.org/x/crypto/bcrypt"
	//"time"
	"bytes"
)

func ChangeAccountImagesAndStatus(eclient *elastic.Client, userID string, image string, status bool, banner string)error{

	err := post.UpdateUser(eclient,userID,"Avatar", image)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"Status", status)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"Banner", banner)
	return err

}

func ChangeContactAndDescription(eclient *elastic.Client, userID string, phone string, phoneVis bool, gender string, genderVis bool, email string, emailVis bool, description string)error{

	err := post.UpdateUser(eclient,userID,"Phone", phone)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"PhoneVis", phoneVis)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"Gender", gender)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"GenderVis", genderVis)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"Email", email)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"EmailVis", emailVis)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"Description", description)
	return err

}

func ChangeFirstAndLastName(eclient *elastic.Client, userID string, first string, last string)error{

	err := post.UpdateUser(eclient,userID,"FirstName", first)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"LastName", last)
	return err

}

func ChangePassword(eclient *elastic.Client, userID string, oldPass []byte, newPass []byte)error{

	usr, err := get.GetUserByID(eclient, userID)
	if (err!= nil){return err}
	if (!bytes.Equal(usr.Password, oldPass)){return errors.New("Invalid old password")}
	err = post.UpdateUser(eclient, userID,"Password", newPass)
	return err

}

func ChangeLocation(eclient *elastic.Client, userID string, country string, state string, stateVis bool, city string, cityVis bool, zip string, zipVis bool)error{

	var newLoc types.LocStruct
	newLoc.Country = country
	newLoc.State = state
	newLoc.StateVis = stateVis
	newLoc.City = city
	newLoc.CityVis = cityVis
	newLoc.Zip = zip
	newLoc.ZipVis = zipVis
	err := post.UpdateUser(eclient, userID,"Location", newLoc)
	return err

}

func ChangeEducation(eclient *elastic.Client, userID string, accType int8, hs string, hsGrad string, uni string, uniGrad string, major []string, minor []string)error{

	err := post.UpdateUser(eclient,userID,"AccType", accType)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"HighSchool", hs)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"HSGradDate", hsGrad)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"University", uni)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"CollegeGradDate", hsGrad)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"Majors", major)
	if (err!= nil){return err}
	err = post.UpdateUser(eclient,userID,"Minors", minor)
	return err
}