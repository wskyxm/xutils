package httpsvr

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

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
