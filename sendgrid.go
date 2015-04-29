package main

import (
	"flag"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"io/ioutil"
	"os"
	"strings"
)

type flagSlice []string

func (f *flagSlice) String() string {
	return fmt.Sprint(*f)
}

func (f *flagSlice) Set(value string) error {
	if *f == nil {
		*f = make(flagSlice, 0)
	}
	*f = append(*f, value)
	return nil
}

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
	ccs := flag.String("cc", "", "CC (comma delimited)")
	bccs := flag.String("bcc", "", "BCC (comma delimited)")
	var attachments flagSlice
	flag.Var(&attachments, "a", "Attachment")
	flag.Parse()

	if *subject == "" || *from == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sg := sendgrid.NewSendGridClient(username, password)
	message := sendgrid.NewMail()
	recipients := strings.Split(flag.Args()[len(flag.Args())-1], ",")
	message.AddTos(recipients)
	message.AddCcs(strings.Split(*ccs, ","))
	message.AddBccs(strings.Split(*bccs, ","))
	
	message.SetSubject(*subject)

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	message.SetText(string(b))
	message.SetFrom(*from)

	for _, attachment := range attachments {
		r, err := os.Open(attachment)
		if err != nil {
			panic(err)
		}
		message.AddAttachment(attachment, r)
	}

	/* Send Message */
	if r := sg.Send(message); r == nil {
		fmt.Println("Email sent!")
	} else {
		panic(r)
	}
}
