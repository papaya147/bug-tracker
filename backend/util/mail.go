package util

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"
)

type SendMailArgs struct {
	From          string
	Password      string
	To            string
	Subject       string
	TemplatePath  string
	TemplateData  map[string]interface{}
	EmailHost     string
	EmailHostPort int
}

func SendMail(args SendMailArgs) {
	body, err := parseTemplate(args.TemplatePath, args.TemplateData)
	if err != nil {
		log.Println(err)
		return
	}

	msg := "From: " + args.From + "\r\n" +
		"To: " + args.To + "\r\n" +
		"Subject: " + args.Subject + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body

	auth := smtp.PlainAuth("", args.From, args.Password, args.EmailHost)
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", args.EmailHost, args.EmailHostPort), auth, args.From, []string{args.To}, []byte(msg)); err != nil {
		log.Println(err)
	}
}

func parseTemplate(templatePath string, data any) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
