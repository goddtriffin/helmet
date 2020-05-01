# <img src="docs/logo.svg" alt="HelmetJS logo" height="64px"/> Helmet

HTTP security middleware for [Go(lang)](https://golang.org/) inspired by [HelmetJS](https://helmetjs.github.io/).

## Quick Start

`go get github.com/MagnusFrater/helmet`

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MagnusFrater/helmet"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I love HelmetJS, I wish there was a Go(lang) equivalent...")
	})

	helmet := helmet.Default()
	http.Handle("/", helmet.Secure(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

## How It Works

Helmet is a collection of 12 smaller middleware functions that set HTTP security response headers. Initializing via `helmet.Default()` will not include all of these middleware functions by default.

| Module                                                                                                     | Default                           |
| ---------------------------------------------------------------------------------------------------------- | --------------------------------- |
| [Content-Security-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP)                           |                                   |
| [X-DNS-Prefetch-Control](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-DNS-Prefetch-Control) | `off`                             |
| [Expect-CT](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect-CT)                           |                                   |
| [Feature-Policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Feature-Policy)                 |                                   |
| [X-Frame-Options](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options)               | `SAMEORIGIN`                      |
| [X-Permitted-Cross-Domain-Policies](https://helmetjs.github.io/docs/crossdomain/)                          |                                   |
| [X-Powered-By](https://helmetjs.github.io/docs/hide-powered-by/)                                           | Removes the `X-Powered-By` header |
