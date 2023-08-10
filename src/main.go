package main

import (
	"net/http"
	"xutils/src/xconf"
	"xutils/src/xerr"
	"xutils/src/xhttp/httpres"
	"xutils/src/xhttp/httpsvr"
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

	var config struct{
		xconf.Configure
		ServerAddr string
	}

	xconf.Load("./config.toml", &config)

	xlog.SetLogLevel(config.LogLevel)
	xlog.SetLogDir(config.LogDir)

	svr := httpsvr.New(httpsvr.Config{AuthCB: auth, Root: "./www"})

	svr.Auth(svr.GET,"/test", test)
	svr.Bind(svr.GET,"/test1", test)

	svr.Run(":8080")
}
