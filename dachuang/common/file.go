package common

import (
	"errors"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
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
