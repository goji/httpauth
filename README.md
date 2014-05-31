# goji/httpauth

httpauth currently provides [HTTP Basic Authentication middleware for](http://tools.ietf.org/html/rfc2617) for [Goji](https://goji.io/), a mimimal web framework for Go.

## Example

httpauth provides a `SimpleBasicAuth` function to get you up and running. Particularly ideal for development servers.

Note that HTTP Basic Authentication credentials are sent over the wire "in the clear" (read: plaintext!) and therefore should not be considered a robust way to secure a HTTP server. If you're after that, you'll need to use SSL/TLS ("HTTPS") at a minimum.

```go

package main

import(
    "net/http"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/middleware"
)

func main() {

    goji.Use(httpauth.SimpleBasicAuth("dave", "somepassword"), middleware.SomeOtherMiddleware)
    // myHandler requires HTTP Basic Auth
    goji.Get("/thing", myHandler)
    goji.Abandon(SimpleBasicAuth)
    // indexHandler does not.
    goji.Get("/", indexHandler)

    goji.Serve()
}
```

If you're looking for a little more control over the process, you can instead pass a `httpauth.AuthOptions` struct to `httpauth.BasicAuth` instead. This allows you to:

* Configure the authentication realm
* Provide your own UnauthorizedHandler (anything that satisfies `http.Handler`) so you can return a better looking 401 page.
* Pass in a custom Validation function that takes the username and password strings and returns a bool, in the event you want to handle specific cases/users.

```go

func main() {

    authOpts := httpauth.AuthOptions{
        Realm: "DevCo",
        User: "dave",
        Password: "plaintext!",
        UnauthorizedHandler: myUnauthorizedHandler,
    }

    goji.Use(BasicAuth(authOpts), myOtherMiddleware)
    goji.Get("/thing", myHandler)

    goji.Serve()
}
```

## Contributing

Send a pull request!
