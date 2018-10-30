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
	SMTPUser     = "ustarttestemail@gmail.com"
	SMTPPass     = "Ust@rt20!8~~"
	Host         = "smtp.gmail.com"
	SendMailPort = 587 //alternatively: 25
	Tags         = "genre=test,genre2=test2"
	SenderEmail  = "ustarttestemail@gmail.com"
)
