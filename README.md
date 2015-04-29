# sendgrid
Send email using Sendgrid using a command line interface similar to the sendmail

# example

```
SENDGRID_USERNAME=username SENDGRID_PASSWORD=password sendgrid -s 'Some subject' -f fromemail@gmail.com toemail@gmail.com < sendgrid.go
```

Send an email to 2 recipients

```
SENDGRID_USERNAME=username SENDGRID_PASSWORD=password sendgrid -s '2 recipients' -f fromemail@gmail.com toemail@gmail.com,toemail2@gmail.com < sendgrid.go
```

Send attachments

(globals implied)

```
sendgrid -s '1 attachment' -f fromemail@gmail.com toemail@gmail.com -a README.md < sendgrid.go
sendgrid -s '2 attachments' -f fromemail@gmail.com toemail@gmail.com -a README.md -a sendgrid.go < sendgrid.go
```

# installation

Download a binary for your environment:
https://github.com/lancecarlson/sendgrid/releases