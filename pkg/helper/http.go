package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
	TraceId string `json:"trace_id"`
}

type PageResult struct {
	List     any   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data any, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
		c.GetString("x-request-id"),
	})
}

func ErrorReturn(msg string, c *gin.Context) {
	c.String(500, msg)
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]any{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

// Excel
// 前端参考 https://juejin.cn/post/6844904014140686349
// https://www.kandaoni.com/news/30927.html
// 导出为 xlsx 文件, filename 可以不包含 xlsx 后缀
//func Excel(titleList []string, dataList [][]any, filename string, c *gin.Context) {
//	xlsx := excelize.NewFile()
//	sheetName := "sheet1"
//
//	for idx, title := range titleList {
//		xlsx.SetCellValue(sheetName, cellName(idx, 1), title)
//	}
//
//	for line, rowData := range dataList {
//		for col, data := range rowData {
//			xlsx.SetCellValue(sheetName, cellName(col, line+2), data)
//		}
//	}
//
//	c.Header("Content-Type", "application/octet-stream")
//	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
//	c.Header("Content-Transfer-Encoding", "binary")
//	xlsx.Write(c.Writer)
//}
//
//var (
//	sheetWords = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
//		"V", "W", "X", "Y", "Z"}
//)
//
//func cellName(col, line int) string {
//	word := sheetWords[col]
//	return fmt.Sprintf("%s%d", word, line)
//}
