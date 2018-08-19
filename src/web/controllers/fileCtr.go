package controllers

import (
	. "../../datamodels"
	. "../../services"
	. "../../tools/http"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func FileUpload(ctx *Context) {
	result := Result{}
	// Get the file from the request.
	file, inf, err := ctx.FormFile("file")
	fileSuffix := path.Ext(inf.Filename)
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
		ctx.JSON(result)
	}
	defer file.Close()
	fname := uuid.Must(uuid.NewV4()).String() + fileSuffix
	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile("./web/views/img/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
		ctx.JSON(result)
	}
	defer out.Close()

	addFile := FileVo{}
	addFile.ID = fname
	addFile.Type = 203
	io.Copy(out, file)

	err2 := addFile.FileInsert()
	if err2 != nil {
		result.Code = 0
		result.Msg = err.Error()
		ctx.JSON(result)
	}

	result.Code = 200
	result.Msg = "成功保存code信息"
	result.Data = fname
	ctx.JSON(result)
}

func FileGet(ctx *Context) {
	fileName := ctx.Params().Get("name")
	f, err := ioutil.ReadFile("./web/views/img/" + fileName) // For read access.
	if err != nil {
		ctx.WriteString("没有文件")
	}
	ctx.WriteString(string(f))
}
