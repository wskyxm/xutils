package httpsvr

import (
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wskyxm/xutils/xlog"
)

type H = gin.H
type IRoutes = gin.IRoutes
type Context = gin.Context
type HandlerFunc = gin.HandlerFunc

type Config struct {
	Root           string
	IndexPage      string
	HandlerUse     []HandlerFunc
	HandlerNoRoute HandlerFunc
	AuthCB         HandlerFunc
	AuthFCB        HandlerFunc
	AuthBCB        HandlerFunc
}

type Server struct {
	config Config
	router *gin.Engine
}

func New(config Config) *Server {
	// 设置GIN日志
	gin.DefaultWriter = io.Discard
	gin.SetMode(gin.ReleaseMode)

	// 创建WEB服务器实例
	svr := &Server{config: config}

	// 初始化默认页
	svr.config.IndexPage = strings.TrimPrefix(svr.config.IndexPage, "/")
	if svr.config.IndexPage == "" {
		svr.config.IndexPage = "index.html"
	}

	// 初始化WEB服务器实例
	svr.router = gin.New()
	svr.router.SetTrustedProxies(nil)
	svr.router.NoRoute(svr.noroute)
	svr.router.Use(append(config.HandlerUse, svr.cors)...)

	// 设置自定义默认路由
	if svr.config.HandlerNoRoute != nil {
		svr.router.NoRoute(svr.config.HandlerNoRoute)
	}

	// 返回实例
	return svr
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func (s *Server) Run(addr string) error {
	xlog.Info("http server startup at %s", addr)
	return s.router.Run(addr)
}
