package contact

import (
	"bytes"
	_ "embed"
	"html/template"
	"log"
	"net/smtp"
)

//go:embed "contact.txt"
var Tmpl string

const host = "mail.martianoids.com"
const port = "25"
const user = "no-reply@abadiaretro.com"
const pass = "La4n8Q#mRjdjyh!Zd*&X"
const to = "sistemas@martianoids.com"

// Contact struct
type Contact struct {
	From    string
	ReplyTo string
	To      string
	Name    string
	Subject string
	Message string
	IP      string
	Agent   string
	Lang    string
}

// New contact struct
func New() *Contact {
	return new(Contact)
}

func (c *Contact) Send() {
	c.From = user
	c.To = to
	t := template.Must(template.New("contact").Parse(Tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, c); err != nil {
		panic(err)
	}

	auth := smtp.PlainAuth("", user, pass, host)
	if err := smtp.SendMail(host+":"+port, auth, user, []string{c.To}, buf.Bytes()); err != nil {
		log.Fatalln("SMTP ERROR: ", err)
	}
}
