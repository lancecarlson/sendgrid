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

# installation

Download a binary for your environment:
https://github.com/lancecarlson/sendgrid/releases