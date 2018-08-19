package controllers

import (
	. "../../datamodels"
	. "../../tools/http"
	"io"
	"io/ioutil"
	"os"
)

func FileUpload(ctx *Context) {
	result := Result{}
	// Get the file from the request.
	file, info, err := ctx.FormFile("file")
	if err != nil {
		result.Code = 0
		result.Msg = err.Error()
		ctx.JSON(result)
	}
	defer file.Close()
	fname := info.Filename
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
	io.Copy(out, file)
	result.Code = 200
	result.Msg = "成功保存code信息"
	result.Data = fname
	ctx.JSON(result)
}

func FileGet(ctx *Context) {
	f, err := ioutil.ReadFile("./web/views/img/典型页面1.png") // For read access.
	if err != nil {
	}
	ctx.WriteString(string(f))
}
