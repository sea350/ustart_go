package globals

const (
	Port = 5002

	//Elastic Client
	//var Eclient, clientErr = elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	//Client URL...
	ClientURL = "http://localhost:9200"

	//SiteURL
	SiteURL = "http://ustart.today"

	//MIME ...
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	Sender     = "ustarttestemail@gmail.com"
	SenderName = "USTART"
	// var serve = "smtp.gmail.com"
	SMTPUser     = "AKIAJNA5EV7IQ5NA6GMQ"
	SMTPPass     = "AlfAq8CoUxc9Vx/FxgRLuYkpgdGuR3ZDCqVM9BoXtDs/"
	Host         = "email-smtp.us-east-1.amazonaws.com"
	SendMailPort = 587 //alternatively: 25
	Tags         = "genre=test,genre2=test2"
)
