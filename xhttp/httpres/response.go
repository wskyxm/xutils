package httpres

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/wskyxm/xutils/xerr"
	"github.com/wskyxm/xutils/xlog"
	"io"
	"net/http"
)

var escapeHTML bool

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (res *ResponseData)String() string {
	data, _ := json.Marshal(res)
	return string(data)
}

func SetEscapeHTML(bool enable) {
	escapeHTML = enable
}

func NewResponseData(code int, err error, data interface{}) *ResponseData {
	err2str := func()string{if err == nil {return "ok"} else {return err.Error()}}
	return &ResponseData{Code: code, Message: err2str(), Data: data}
}

func NewSuccess() *ResponseData {
	return NewResponseData(http.StatusOK, nil, nil)
}

func respone(c *gin.Context, code int, err error, data interface{}, level xlog.Severity) {
	// 参数检查
	if err == nil {err = xerr.NoError; code = http.StatusOK}
	
	// 数据编码
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	enc.SetEscapeHTML(escapeHTML)
	enc.Encode(ResponseData{Code: code, Message: err.Error(), Data: data})

	// 发送数据
	c.Status(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.Write(buffer.Bytes())
}

func Error(c *gin.Context, err error) {
	switch err.(type) {
	case *xerr.Error: respone(c, err.(*xerr.Error).Code, err, nil, xlog.ErrorLog)
	default: respone(c, http.StatusInternalServerError, err, nil, xlog.ErrorLog)
	}
}

func ErrorWithData(c *gin.Context, err error, data interface{}) {
	switch err.(type) {
	case *xerr.Error: respone(c, err.(*xerr.Error).Code, err, data, xlog.ErrorLog)
	default: respone(c, http.StatusInternalServerError, err, data, xlog.ErrorLog)
	}
}

func Respone(c *gin.Context, err error, data interface{}) {
	if err != nil {
		ErrorWithData(c, err, data)} else {
		Success(c, data)}
}

func Success(c *gin.Context, data interface{}) {
	respone(c, http.StatusOK, nil, data, xlog.InfoLog)
}

func GetResponeData(body io.ReadCloser) *ResponseData {
	resp := &ResponseData{}
	GetBody(body, resp); return resp
}

func Request(method, uri string, header map[string]string, body []byte, obj any) ([]byte, error) {
	// 创建请求
	req, err := http.NewRequest(method, uri, bytes.NewReader(body))
	if err != nil {return nil, err}

	// 设置请求头
	for key, val := range header {req.Header.Set(key, val)}

	// 发起请求
	cli := &http.Client{}
	rep := (*http.Response)(nil)
	if rep, err = cli.Do(req); err != nil {return nil, err}

	// 读取返回数据
	return GetBody(rep.Body, obj)
}

func GetBody(body io.ReadCloser, obj any) ([]byte, error) {
	// 退出时关闭
	defer body.Close()

	// 读取数据
	data, err := io.ReadAll(body)
	if err != nil {return nil, err}

	// 解析数据
	if obj != nil {err = json.Unmarshal(data, obj)}
	return data, err
}
