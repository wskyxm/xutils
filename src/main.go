package main

import (
	"fmt"
	"net/http"
	"xutils/src/xcache"
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

	xcache.Initialize("", "")
	xcache.HSet("test1", map[string]string{"abc": "123"})
	xcache.Set("test2", "456", 0)
	fmt.Println(xcache.HGet("test1", "abc"))
	fmt.Println(xcache.HGetAll("test1"))
	fmt.Println(xcache.Get("test2"))

	svr.Run(":8080")
}
