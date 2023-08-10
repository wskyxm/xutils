package httpsvr

import (
	"github.com/gin-gonic/gin"
	"xutils/src/xlog"
)

type MethodFunc func(relativePath string, handlers ...HandlerFunc)IRoutes

func (s *Server)auth(method MethodFunc, path string, fn HandlerFunc, handler ...gin.HandlerFunc) {
	// 参数检查
	if fn == nil {xlog.Panic("auth callback is nil")}

	// 绑定路由处理函数
	method(path, append([]gin.HandlerFunc{fn}, handler...)...)
}

func (s *Server)Bind(method MethodFunc, path string, handler ...gin.HandlerFunc) {
	method(path, handler...)
}

func (s *Server)Auth(method MethodFunc, path string, handler ...gin.HandlerFunc) {
	s.auth(method, path, s.config.AuthCB, handler...)
}

func (s *Server)AuthF(method MethodFunc, path string, handler ...gin.HandlerFunc) {
	s.auth(method, path, s.config.AuthFCB, handler...)
}

func (s *Server)AuthB(method MethodFunc, path string, handler ...gin.HandlerFunc) {
	s.auth(method, path, s.config.AuthBCB, handler...)
}

