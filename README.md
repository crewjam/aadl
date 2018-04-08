# AADL Checkouts

The old Ann Arbor District Library site had an XML feed of your 
checked out materials, but the new one doesn't appear to. This
little program prints a text list of everything you have checked
out.

Install:

```bash
go build -o ~/bin/aadl aadl.go
```

Usage:

```bash
aadl -username YourUsername -password 12345
```

I have a cron that runs every Saturday and shoots the output of this
to a receipt printer I got on ebay.