package post

import (
	elastic "gopkg.in/olivere/elastic.v5"
)

//UpdateSignUpWarningByIP ...
func UpdateSignUpWarningByIP(eclient elastic.Client, addressIP string, field string, newContent interface{}) {
	//code
}
