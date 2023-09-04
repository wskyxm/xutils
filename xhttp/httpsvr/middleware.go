package httpsvr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server)cors(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	context.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Referer, token, User-Agent, Accept, Authorization")
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Original-Name, Content-Disposition")
	context.Header("Access-Control-Allow-Credentials", "true")
	context.Header("Cross-Origin-Embedder-Policy", "require-corp")
	context.Header("Cross-Origin-Opener-Policy", "same-origin")

	// 检查请求方式
	if context.Request.Method != "OPTIONS" {context.Next()
	} else {context.AbortWithStatus(http.StatusNoContent)}
}

func (s *Server)noroute(c *gin.Context) {
	// 参数检查
	if s.wwwdir == nil {return}

	// 静态文件服务
	c.FileFromFS(c.Request.URL.Path, s.wwwdir)
}
