package config

import (
	"fmt"
	"strconv"
	"github.com/go-ini/ini"
)

var (
	iniFile *ini.File
)

func init(){
	file,err := ini.Load("config/config.ini")
	if err != nil{
		fmt.Println("[+]配置文件加载失败")
		return
	}
	iniFile = file
}

func GetSection(sectionName string) *ini.Section {
	section,err := iniFile.GetSection(sectionName)
	if err != nil{
		fmt.Println("[*]没有找到对应配置字段")
		return nil
	}
	return section
}

//Get string value from section && key
func GetValue(sectionName string,key string) string {
	var val string
	section := GetSection(sectionName)
	if section != nil{
		val = section.Key(key).Value()
	}
	return val;
}

//Get bool value from section && key
func GetBool(sectionName string,key string) bool {
	val := GetValue(sectionName,key)
	value,err := strconv.ParseBool(val)
	if err != nil{
		fmt.Println(`[-]String Convert to Bool value Failed!`)
	}
	return value
}

//Get int value from section && key
func GetInt(sectionName string,key string) int {
	val := GetValue(sectionName,key)
	value,err:= strconv.Atoi(val)
	if err != nil{
		fmt.Println(`[-]String Convert to Int value Failed!`)
	}
	return value
}