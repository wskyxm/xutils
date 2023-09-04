package httpsvr

import (
	"github.com/gin-gonic/gin"
	"github.com/wskyxm/xutils/xlog"
	"io"
	"net/http"
)

type H = gin.H
type IRoutes = gin.IRoutes
type Context = gin.Context
type HandlerFunc = gin.HandlerFunc

type Config struct {
	Root           string
	HandlerNoRoute []HandlerFunc
	HandlerUse     []HandlerFunc
	AuthCB         HandlerFunc
	AuthFCB        HandlerFunc
	AuthBCB        HandlerFunc
}

type Server struct {
	wwwdir http.FileSystem
	config Config
	router *gin.Engine
}

func New(config Config) *Server {
	// 设置GIN日志
	gin.DefaultWriter = io.Discard
	gin.SetMode(gin.ReleaseMode)

	// 创建WEB服务器实例
	svr := &Server{config: config}
	if config.Root != "" {svr.wwwdir = http.Dir(config.Root)}

	// 初始化WEB服务器实例
	svr.router = gin.New()
	svr.router.SetTrustedProxies(nil)
	svr.router.NoRoute(append(config.HandlerNoRoute, svr.noroute)...)
	svr.router.Use(append(config.HandlerUse, svr.cors)...)

	// 返回实例
	return svr
}

func (s *Server)Run(addr string) {
	xlog.Info("http server startup at %s", addr)
	s.router.Run(addr)
}