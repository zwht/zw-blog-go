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
	"strconv"
	"strings"
)

func FileUpload(ctx *Context) {
	result := Result{}
	// Get the file from the request.
	file, inf, err := ctx.FormFile("file")
	fileSuffix := path.Ext(inf.Filename)
	fileType := ctx.FormValue("fileType")
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
	addFile.Type, _ = strconv.Atoi(fileType)
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
func FileGetList(ctx *Context) {
	pageSize, _ := strconv.Atoi(ctx.Params().Get("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Params().Get("pageNum"))
	var fileSearchVo FileSearchVo
	ctx.ReadJSON(&fileSearchVo)
	count, _ := FileSelectCount(fileSearchVo)
	files, err := FileSelectList(pageSize, pageNum, fileSearchVo)
	result := Result{}
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
	} else {
		resultPage := ResultPage{}
		resultPage.TotalCount = count
		resultPage.PageSize = pageSize
		resultPage.PageNum = pageNum
		resultPage.PageData = files

		result.Code = 200
		result.Msg = "成功获取列表信息"
		result.Data = resultPage
	}

	ctx.JSON(result)
}

func FileDeleteById(ctx *Context) {
	id := ctx.Params().Get("id")
	result := Result{}

	del := os.Remove("./web/views/img/" + id)
	if del != nil && strings.Contains(del.Error(), "no such file or directory") != true {
		result.Code = 0
		result.Msg = del.Error()
		ctx.JSON(result)
	} else {
		err := FileDelete(id)
		if err != nil {
			result.Code = 0
			result.Msg = err.Error()
			ctx.JSON(result)
		}
		result.Code = 200
		result.Msg = "成功删除"
		ctx.JSON(result)
	}
}
