package contact

import (
	"bytes"
	_ "embed"
	"html/template"
	"log"
	"net/smtp"

	"git.martianoids.com/queru/retroserver/internal/cfg"
)

//go:embed "contact.txt"
var Tmpl string

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
	c.From = cfg.SMTPUser
	c.To = cfg.SMTPTo
	t := template.Must(template.New("contact").Parse(Tmpl))
	var buf bytes.Buffer
	if err := t.Execute(&buf, c); err != nil {
		panic(err)
	}

	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
	if err := smtp.SendMail(cfg.SMTPHost+":"+cfg.SMTPPort, auth, cfg.SMTPUser, []string{c.To},
		buf.Bytes()); err != nil {
		log.Fatalln("SMTP ERROR: ", err)
	}
}
