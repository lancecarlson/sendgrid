# sendgrid-cli
Send email using sendgrid using a CLI similar to the sendmail command

# example

```
SENDGRID_USERNAME=username SENDGRID_PASSWORD=password sendgrid-cli -s 'Some subject' -f fromemail@gmail.com toemail@gmail.com < sendgrid.go
```