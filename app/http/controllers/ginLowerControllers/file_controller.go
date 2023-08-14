package ginLowerControllers

import (
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"thh/app/bundles/logging"
	"thh/app/http/controllers/component"
	"time"

	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
)

func getPicB() []byte {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}

	buffer := bytes.NewBuffer(nil)
	err := dc.EncodePNG(buffer)
	if err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}

func GinShowPic(c *gin.Context) {
	c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "0")
	c.Data(http.StatusOK, "image/png", getPicB())
}

func GinUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("获取数据失败")
		ctx.JSON(http.StatusOK, component.DataMap{
			"code":    1,
			"message": "获取数据失败" + err.Error(),
		})
		return
	}

	fmt.Println("接收的数据", file.Filename)
	//获取文件名称
	fmt.Println(file.Filename)
	//文件大小
	fmt.Println(file.Size)
	//获取文件的后缀名
	fileExt := path.Ext(file.Filename)
	fmt.Println(fileExt)
	//根据当前时间鹾生成一个新的文件名
	fileNameInt := time.Now().Unix()
	fileNameStr := cast.ToString(fileNameInt)
	//新的文件名
	fileName := fileNameStr + fileExt
	//保存上传文件
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join("./storage/upload", folderName)
	//使用 MkdirAll 会创建多层级目录
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		logging.Error("创建目录失败:" + err.Error())
		ctx.JSON(http.StatusInternalServerError, component.DataMap{
			"code":    1,
			"message": "保存错误" + err.Error(),
		})
		return
	}
	filePath := filepath.Join(folderPath, "/", fileName)
	err = ctx.SaveUploadedFile(file, filePath)
	msg := "SUCCESS"
	if err != nil {
		msg = err.Error()
	}
	ctx.JSON(http.StatusOK, component.DataMap{
		"code": component.SUCCESS,
		"result": component.DataMap{
			"filePath": filePath,
		},
		"message": msg,
	})
}
