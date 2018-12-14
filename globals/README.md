# How to define envVariable.go

## Create a new [Go] file using the name "envVariables"

Copy the following code into that new file and fill in any feild marked in the comments

```
package globals

const (
	//Port is the current operative port
	Port = "5002"

	//ClientURL is the elastic client url
	ClientURL = "" //FILL THIS IN

	//SiteURL is the site domain
	SiteURL = "http://ustart.today"

	//HTMLPATH is how to get to the html filed from the working directory
	HTMLPATH = "../ustart_front/"

	//MIME is a pokemon
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	//Sender ...
	Sender = "ustarttestemail@gmail.com"
	//SenderName ...
	SenderName = "U•START"
	//SMTPUser ...
	SMTPUser = "" //FILL THIS IN
	//SMTPPass ..
	SMTPPass = "" //FILL THIS IN
	//Host ...
	Host = "email-smtp.us-east-1.amazonaws.com"
	//SendMailPort ...
	SendMailPort = 587 //alternatively: 25
	//EmailTags ...
	EmailTags = "genre=test,genre2=test2"
	//SenderEmail ...
	SenderEmail = "ustarttestemail@gmail.com"

	//S3 GLOBALS

	//S3Region The geographic AWS region the S3 instance resides
	S3Region = "us-east-1"

	//The credentials needed to acess S3 instance

	//S3CredID ID
	S3CredID = "" //FILL THIS IN
	//S3CredSecret Secret
	S3CredSecret = ""//FILL THIS IN
	//S3CredToken Token
	S3CredToken = ""
	//S3BucketName The name of the bucket itself
	S3BucketName = "ustart-bucket"
)
```