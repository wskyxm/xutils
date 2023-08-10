package httpsvr

import (
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"xutils/src/xerr"
	"xutils/src/xhttp/httpres"
)

func HttpFirstFile(c *gin.Context) (*multipart.FileHeader, error) {
	// 解析表单
	form, err := c.MultipartForm()
	if err != nil {return nil, err}

	// 遍历文件列表
	for _, files := range form.File {
		if len(files) == 0 {continue} else {return files[0], nil}
	}

	// 没有找到上传的文件
	return nil, xerr.InvalidParameter
}

func HttpSaveFile(body io.ReadCloser, length int64, path string) error {
	// 初始化缓冲区
	buff := make([]byte, 1024)
	file := (*os.File)(nil)
	werr := error(nil)
	size := int64(0)
	read := 0

	// 创建文件
	if file, werr = os.Create(path); werr != nil {return werr}

	// 返回时关闭
	defer func() {body.Close(); file.Close()}()

	// 保存文件
	for {
		read, werr = body.Read(buff)
		if read > 0 {if _, werr = file.Write(buff[:read]); werr == nil {size += int64(read)}}
		if werr != nil {break}
	}

	// 返回结果
	if size != length {return werr} else {return nil}
}

func HttpSaveFirstFile(c *gin.Context, path string) (string, error) {
	// 获取上传文件信息
	filehdr, err := HttpFirstFile(c)
	if err != nil {return "", err}

	// 打开上传的文件
	file, err := filehdr.Open()
	if err != nil {
		httpres.Error(c, err); return "", err}

	// 保存文件到本地
	return filehdr.Filename, HttpSaveFile(file, filehdr.Size, path)
}
