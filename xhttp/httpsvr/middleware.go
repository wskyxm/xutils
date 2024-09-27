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
	// 获取请求的文件路径
	path := filepath.Join(s.config.Root, c.Request.URL.Path)

	// 如果是目录，就增加默认页
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		path = strings.TrimSuffix(path, "/") + "/" + s.config.IndexPage
	}

	// 检查文件是否存在
	if _, err := os.Stat(path); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// 返回文件内容
	http.ServeFile(c.Writer, c.Request, path)
}
