package main

import (
	"github.com/wskyxm/xutils/xconf"
	"github.com/wskyxm/xutils/xerr"
	"github.com/wskyxm/xutils/xhttp/httpres"
	"github.com/wskyxm/xutils/xhttp/httpsvr"
	"github.com/wskyxm/xutils/xlog"
	"net/http"
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
