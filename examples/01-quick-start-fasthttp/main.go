package main

import (
	"github.com/fasthttp/router"
	"github.com/goddtriffin/helmet"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()

	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("I love HelmetJS, I just wish there was a Go(lang) equivalent...")
	})

	h := helmet.Default()

	httpServer := fasthttp.Server{
		Handler: h.SecureFastHTTP(r.Handler),
	}

	httpServer.ListenAndServe(":8080")
}
