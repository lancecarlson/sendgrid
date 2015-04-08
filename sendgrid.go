package main

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"github.com/sendgrid/sendgrid-go"
)

func main() {
	var username = os.Getenv("SENDGRID_USERNAME")
	var password = os.Getenv("SENDGRID_PASSWORD")

	if username == "" {
		panic("SENDGRID_USERNAME environment variable not set")
	}

	if password == "" {
		panic("SENDGRID_PASSWORD environment variable not set")
	}

	var subject = flag.String("s", "", "Subject")
	var from = flag.String("f", "", "From")
	flag.Parse()

	if *subject == "" || *from == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sg := sendgrid.NewSendGridClient(username, password)
	message := sendgrid.NewMail()
	for _, recipient := range flag.Args() {
		message.AddTo(recipient)
	}
//	message.AddToName("Yamil Asusta")
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
