# goji/httpauth [![GoDoc](https://godoc.org/github.com/goji/httpauth?status.png)](https://godoc.org/github.com/goji/httpauth)

httpauth currently provides [HTTP Basic Authentication middleware for](http://tools.ietf.org/html/rfc2617) for Go. 

Note that httpauth is completely compatible with [Goji](https://goji.io/), a mimimal web framework for Go, but as it satisfies http.Handler it can be used beyond Goji itself. 
## Example

httpauth provides a `SimpleBasicAuth` function to get you up and running. Particularly ideal for development servers.

Note that HTTP Basic Authentication credentials are sent over the wire "in the clear" (read: plaintext!) and therefore should not be considered a robust way to secure a HTTP server. If you're after that, you'll need to use SSL/TLS ("HTTPS") at a minimum.

### Goji

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

    goji.Serve()
}
```

If you're looking for a little more control over the process, you can instead pass a `httpauth.AuthOptions` struct to `httpauth.BasicAuth` instead. This allows you to:

* Configure the authentication realm
* Provide your own UnauthorizedHandler (anything that satisfies `http.Handler`) so you can return a better looking 401 page.

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

### net/http

If you're using vanilla net/http or another http.Handler compatible framework:

```go
package main

import(
	"net/http"
	"github.com/goji/httpauth"
)

func main() {
	http.Handle("/", httpauth.SimpleBasicAuth("dave", "somepassword")(http.HandlerFunc(hello)))
	http.ListenAndServe(":7000", nil)
}
```

## Contributing

Send a pull request! Note that features on the (informal) roadmap include HTTP Digest Auth and the potential for supplying your own user/password comparison function.
