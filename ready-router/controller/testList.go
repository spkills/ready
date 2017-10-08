package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func TestListGetHandler(ctx *fasthttp.RequestCtx) {

	name := "test/list"

	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, name)

}

func TestListPostHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func TestListPutHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func TestListDeleteHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusNotImplemented)
	fmt.Fprintf(ctx, "NotImplemented")
}

func TestListHeadHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "")
}
