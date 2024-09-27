package httpsvr

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Server) cors(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", context.GetHeader("Origin"))
	context.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	context.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Referer, token, User-Agent, Accept, Authorization")
	context.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Original-Name, Content-Disposition")
	context.Header("Access-Control-Allow-Credentials", "true")
	context.Header("Cross-Origin-Embedder-Policy", "require-corp")
	context.Header("Cross-Origin-Opener-Policy", "same-origin")

	// 检查请求方式
	if context.Request.Method != "OPTIONS" {
		context.Next()
	} else {
		context.AbortWithStatus(http.StatusNoContent)
	}
}

func (s *Server) noroute(c *gin.Context) {
	// 参数检查
	if s.wwwdir == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// 如果是目录，就增加默认页
	if stat, err := os.Stat(filepath.Join(s.config.Root, c.Request.URL.Path)); err == nil && stat.IsDir() {
		c.Request.URL.Path = strings.TrimSuffix(c.Request.URL.Path, "/") + "/" + s.config.IndexPage
	}

	// 检查文件是否存在
	if _, err := os.Stat(filepath.Join(s.config.Root, c.Request.URL.Path)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// 静态文件服务
	c.FileFromFS(c.Request.URL.Path, s.wwwdir)
}
