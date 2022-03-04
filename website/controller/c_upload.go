package controller

import (
	"errors"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SuffixTips = "图片后缀必须为jpg,jpeg,png"
)

var ValidSuffix = map[string]bool{
	"jpg":  true,
	"png":  true,
	"jpeg": true,
}

func GetFilePath(f *multipart.FileHeader) (file_path string, url_path string, err error) {
	tmp := strings.Split(f.Filename, ".")
	suffix := tmp[len(tmp)-1]
	_, ok := ValidSuffix[suffix]
	if !ok {
		return file_path, url_path, errors.New(SuffixTips)
	}
	new_filename := strconv.FormatInt(time.Now().Unix(), 10) + "." + suffix
	file_path = "./view/static/photo/" + new_filename
	url_path = "/static/photo/" + new_filename
	return file_path, url_path, nil
}


//上传图片
func Upload(c *gin.Context){
	header,err := c.FormFile("image")
	if err != nil{
		ErrorResp(c,"上传错误")
		return
	}
	file_path,file_url,err := GetFilePath(header)
	if err != nil{
		ErrorResp(c,"上传错误")
		return
	}
	err = c.SaveUploadedFile(header,file_path)
	if err != nil{
		ErrorResp(c,"上传错误")
		return
	}
	SuccessResp(c,file_url)

}