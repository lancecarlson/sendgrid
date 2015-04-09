package main

import (
	"flag"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	username := os.Getenv("SENDGRID_USERNAME")
	password := os.Getenv("SENDGRID_PASSWORD")

	if username == "" {
		panic("SENDGRID_USERNAME environment variable not set")
	}

	if password == "" {
		panic("SENDGRID_PASSWORD environment variable not set")
	}

	subject := flag.String("s", "", "Subject")
	from := flag.String("f", "", "From")
	flag.Parse()

	if *subject == "" || *from == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sg := sendgrid.NewSendGridClient(username, password)
	message := sendgrid.NewMail()
	recipients := strings.Split(flag.Args()[len(flag.Args())-1], ",")
	for _, recipient := range recipients {
		message.AddTo(recipient)
	}

	message.SetSubject(*subject)

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	message.SetText(string(b))

	message.SetFrom(*from)
	if r := sg.Send(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		panic(r)
	}
}
