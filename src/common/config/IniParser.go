package config

import (
	"fmt"
)

type IniParser struct {
	filePath string
}

func NewIniParser(filePath string) *IniParser {
	return &IniParser{
		filePath: filePath,
	}
}

func (i *IniParser) Parse() error {
	_, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		return nil
	}
	return nil
}
