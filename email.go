package main

import (
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

const (
	smtpServer = "smtp.gmail.com"
	smtpPort   = 587
)

func sendEmail(config *config, body string) error {
	d := gomail.NewDialer(
		smtpServer,
		smtpPort,
		config.From,
		config.Password,
	)
	sender, err := d.Dial()
	if err != nil {
		return err
	}
	defer sender.Close()
	msg := gomail.NewMessage()
	msg.SetHeader("To", config.EmailAddresses...)
	msg.SetHeader("From", config.From)
	msg.SetHeader("Subject", "This weeks meal plan")
	msg.SetBody("text/html", body)

	err = sender.Send(
		config.From,
		config.EmailAddresses,
		msg,
	)
	if err != nil {
		return err
	}
	return nil
}

func makeBody(config *config, recipes []string) string {
	recipeWrapper := func(recipeFile string) string {
		name := strings.Split(recipeFile, ".")
		return fmt.Sprintf(`<a href="%s/%s">%s</a>`, config.Prefix, recipeFile, name[0])
	}
	output := "<html>\n<body>\n"
	body := "This weeks recipes:<br />\n<br />\n"
	for _, recipe := range recipes {
		body = body + recipeWrapper(recipe) + "<br />\n"
	}
	body = body + "<br />Enjoy!<br />"
	output = output + body
	output = output + "\n</body>\n</html>"
	return output
}
