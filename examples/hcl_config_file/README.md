
```bash
> $ go run main.go -h

Usage
  -domain string
		
  -mongodb_host string
		 (default "example.com")
  -mongodb_port int
		 (default 999)
  -user_name string
		
  -user_passwd string
		
Environment variables:
 $EXAMPLE_DOMAIN string

 $EXAMPLE_USER_NAME string

 $EXAMPLE_USER_PASSWD string

  $EXAMPLE_MONGODB_HOST string
	 (default "example.com")
  $EXAMPLE_MONGODB_PORT int
	 (default "999")

Config file "config.hcl":

 '=' BEFORE '{' IS OPTIONAL

"Domain" = "example.com"

"debug" = true

"mongodb" = {
  "Host" = "myhost"

  "Port" = 9090
}

"user" = {
  "Name" = ""

  "Password" = ""
}
exit status 2
```
