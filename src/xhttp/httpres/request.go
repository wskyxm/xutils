package httpres

import (
	"github.com/gin-gonic/gin"
	"xutils/src/xvalue"
)

type PageInfo struct {
	Page	int
	Size	int
	Total	int64
}

func GetPageInfo(c *gin.Context) PageInfo {
	// 解析请求参数
	info := PageInfo{}
	info.Page = xvalue.S2I(c.Query("page"))
	info.Size = xvalue.S2I(c.Query("size"))

	// 参数检查
	if info.Size <= 0 {info.Size = 10}
	if info.Page <= 1 {info.Page = 1}
	return info
}

func (s PageInfo)Offset() int {
	return (s.Page - 1) * s.Size
}

func (s PageInfo)Pages() int {
	// 参数检查
	if s.Size <= 0 {return int(s.Total)}

	// 返回总页数
	pages := int(s.Total / int64(s.Size))
	if (s.Total % int64(s.Size)) > 0 {pages++}
	return pages
}

func (s PageInfo)Data(items any) gin.H {
	return gin.H{"total": s.Total, "pages": s.Pages(), "items": items}
}