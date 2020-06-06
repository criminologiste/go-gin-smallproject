package v1

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/extrame/xls"
	"github.com/gin-gonic/gin"
	"go-gin-smallproject/models"
	"go-gin-smallproject/pkg/e"
	"log"
	"math/rand"
	"net/http"
	"path"
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
	// 插入数据库的sql
	sql := "INSERT INTO `blog_msg` (`name`,`age`,`address`) VALUES "
	var buf bytes.Buffer
	var sql_str string
	buf.WriteString(sql)
	str_bt := []string{"姓名", "年龄", "家庭住址"}
	//判断上传文件的格式
	if path.Ext(file.Filename) == ".xlsx" {
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
		num_int := 1
		var key0, key1, key2 int
		//  判断表头
		for key, row := range rows {
			if num_int == 1 {
				for k, v := range row {
					if str_bt[0] == v {
						key0 = k
					}
					if str_bt[1] == v {
						key1 = k
					}
					if str_bt[2] == v {
						key2 = k
					}
				}
			}
			// 不插入标题
			if key > 0 {
				if len(rows)-1 == key {
					//最后一条数据 以分号结尾
					sql_str = fmt.Sprintf("('%s','%s','%s');", row[key0], row[key1], row[key2])
					buf.WriteString(sql_str)
				} else if num_int == 3000 {
					sql_str = fmt.Sprintf("('%s','%s','%s');", row[key0], row[key1], row[key2])
					buf.WriteString(sql_str)
					models.InsertExel(buf.String())
					buf.Next(len(buf.String()))
					buf.WriteString(sql)
					code = e.SUCCESS
					num_int = 1
				} else {
					sql_str = fmt.Sprintf("('%s','%s','%s'),", row[key0], row[key1], row[key2])
					buf.WriteString(sql_str)
				}
			}
			num_int++
		}
		if models.InsertExel(buf.String()) {
			code = e.SUCCESS
		}
	} else if path.Ext(file.Filename) == ".xls" {
		var key0, key1, key2 int
		xlFile, err := xls.Open(dst, "utf-8")
		if err != nil {
			log.Fatal(err)
		}
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			num_int := 1
			col1 := sheet1.Row(0).Col(0)
			col2 := sheet1.Row(0).Col(0)
			col3 := sheet1.Row(0).Col(0)
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				col1 = row1.Col(0)
				col2 = row1.Col(1)
				col3 = row1.Col(2)
				row := []string{col1, col2, col3}
				if num_int == 1 {
					for k, v := range row {
						if str_bt[0] == v {
							key0 = k
						}
						if str_bt[1] == v {
							key1 = k
						}
						if str_bt[2] == v {
							key2 = k
						}
					}
				}
				if i > 0 {
					row1 := sheet1.Row(i)
					col1 = row1.Col(key0)
					col2 = row1.Col(key1)
					col3 = row1.Col(key2)
					if int(sheet1.MaxRow) == i {
						//最后一条数据 以分号结尾
						sql_str = fmt.Sprintf("('%s','%s','%s');", col1, col2, col3)
						buf.WriteString(sql_str)
					} else if num_int == 3000 {
						sql_str = fmt.Sprintf("('%s','%s','%s');", col1, col2, col3)
						buf.WriteString(sql_str)
						models.InsertExel(buf.String())
						buf.Next(len(buf.String()))
						buf.WriteString(sql)
						code = e.SUCCESS
						num_int = 1
					} else {
						sql_str = fmt.Sprintf("('%s','%s','%s'),", col1, col2, col3)
						buf.WriteString(sql_str)
					}
				}
				num_int++
			}
		}
		if models.InsertExel(buf.String()) {
			code = e.SUCCESS
		}
	} else {
		code = e.ERROR_EXEL_LOAD
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
