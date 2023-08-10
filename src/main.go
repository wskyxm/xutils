package main

import (
	"net/http"
	"xutils/src/web/httpres"
	"xutils/src/web/httpsvr"
	"xutils/src/xerr"
	"xutils/src/xlog"
)

func auth(c *httpsvr.Context) {
	httpres.Error(c, xerr.AuthenticationFailed)
	c.AbortWithStatus(http.StatusForbidden)
}

func test(c *httpsvr.Context) {
	c.JSON(http.StatusOK, httpsvr.H{"hhhhhh": "vvvvvvv"})
}

func main() {
	xlog.SetLogLevel(xlog.SInfoLog)
	xlog.SetLogDir("")

	svr := httpsvr.New(httpsvr.Config{AuthCB: auth, Root: "./www"})

	svr.Auth(svr.GET,"/test", test)
	svr.Bind(svr.GET,"/test1", test)

	svr.Run(":8080")
}
