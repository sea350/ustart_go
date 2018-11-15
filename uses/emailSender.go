package uses

import (
	"bytes"
	"fmt"
	"net/smtp"

	// "fmt"
	"html/template"
	"log"

	"github.com/sea350/ustart_go/globals"
	// "net/smtp"
)

//Request ...
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

//MIME ...
const MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

//NewRequest ...
func NewRequest(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

func (r *Request) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

func (r *Request) sendMail() bool {
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + globals.MIME + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", globals.Host, globals.SendMailPort)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth(globals.SMTPUser, globals.SenderEmail, globals.SMTPPass, globals.Host), globals.SenderEmail, r.to, []byte(body)); err != nil {
		// log.Println(globals.Host, globals.SenderEmail, globals.SendMailPort, globals.SMTPUser)
		log.Println(err)
		return false
	}
	return true
}

// func (r *Request) sendMail() bool {
// 	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
// 	// SMTP := fmt.Sprintf("%s:%d", Host, Port) //587)

// 	Recipient := r.to[0]
// 	Subject := r.subject

// 	m := gomail.NewMessage()
// 	m.SetBody("text/html", body)
// 	// m.AddAlternative("text/plain", body)
// 	m.SetHeaders(map[string][]string{
// 		"From":    {m.FormatAddress(globals.Sender, globals.SenderName)},
// 		"To":      {Recipient},
// 		"Subject": {Subject},
// 		// "X-SES-Configuration-SET": {ConfigSet},
// 		"X-SES_MESSAGE-TAGS": {globals.EmailTags},
// 	})

// 	d := gomail.NewPlainDialer(globals.Host, globals.SendMailPort, globals.SMTPUser, globals.SMTPPass)

// 	// Display an error message if something goes wrong; otherwise,
// 	// display a message confirming that the message was sent.
// 	if err := d.DialAndSend(m); err != nil {
// 		log.Println(err)
// 		return false
// 	}

// 	log.Println("Email sent!")
// 	return true

// 	// if err := smtp.SendMail(SMTP, smtp.PlainAuth("", "ustarttestemail@gmail.com", "Ust@rt20!8~~", host), "ustarttestemail@gmail.com", r.to, []byte(body)); err != nil {
// 	// 	return false
// 	// }

// }

//Send ...
func (r *Request) Send(templateName string, items interface{}) {
	err := r.parseTemplate(templateName, items)
	if err != nil {
		log.Fatal(err)
	}
	if ok := r.sendMail(); ok {
		log.Printf("Email has been sent to %s\n", r.to)
	} else {
		log.Printf("Failed to send the email to %s\n", r.to)
	}
}
