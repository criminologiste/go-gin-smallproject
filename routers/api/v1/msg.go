package v1

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"go-gin-smallproject/models"
	"go-gin-smallproject/pkg/e"
	"log"
	"math/rand"
	"net/http"
	"path"
	"reflect"
	"strconv"
	"time"
)

// RandString 生成随机字符串
func RandString(len int) string {
	var r *rand.Rand
	r = rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func UpdataExel(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// 上传文件到指定的路径
	dst := path.Join("./static", file.Filename)
	_ = c.SaveUploadedFile(file, dst)
	code := e.INVALID_PARAMS
	// 读取exel 文档
	f, err := excelize.OpenFile(dst)
	if err != nil {
		code = e.ERROR_EXEL_LOAD
		log.Printf(" err.message: %s", err.Error())
	}
	// 获取工作表中指定单元格的值
	cell := f.GetCellValue("Sheet1", "A2")
	if cell == "" {
		code = e.ERROR_EXEL_LOAD
		log.Printf("未填写数据!")
	}
	rows := f.GetRows("Sheet1")
	fmt.Println(reflect.TypeOf(rows).String())
	sql := "INSERT INTO `blog_msg` (`name`,`age`,`address`) VALUES "
	for key, row := range rows {
		// 不插入标题
		if key > 0 {
			if len(rows)-1 == key {
				//最后一条数据 以分号结尾
				sql += fmt.Sprintf("('%s','%s','%s');", row[0], row[1], row[2])
			} else {
				sql += fmt.Sprintf("('%s','%s','%s'),", row[0], row[1], row[2])
			}
		}
	}
	if models.InsertExel(sql) {
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

func DownExel(c *gin.Context) {
	code := e.INVALID_PARAMS
	data := models.GetExelDate()
	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 设置表头
	f.SetCellValue("Sheet1", "A1", "姓名")
	f.SetCellValue("Sheet1", "B1", "年龄")
	f.SetCellValue("Sheet1", "C1", "家庭住址")
	// 写入数据
	intnum := 2
	for _, value := range data {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(intnum), value.Name)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(intnum), value.Age)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(intnum), value.Address)
		intnum = intnum + 1
	}
	unix1 := RandString(11)
	filepath := "./static/" + unix1 + ".xlsx"
	if err := f.SaveAs(filepath); err != nil {
		code = e.ERROR_EXEL_LOAD
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
	} else {
		c.File(filepath)
	}

}
