package parser

import (
	"fmt"
	yamlConfig "github.com/olebedev/config"
)

type YamlParser struct {
	filePath string
	file     *yamlConfig.Config
}

func NewYamlParser(filePath string) *YamlParser {
	return &YamlParser{
		filePath: filePath,
	}
}

func (i *YamlParser) Parse() error {
	file, err := yamlConfig.ParseYamlFile(i.filePath)
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
		return nil
	}
	fmt.Println("file", file)
	i.file = file
	return nil
}

func (i *YamlParser) GetString(sec string, key string) string {
	fmt.Println("GetString....")
	fmt.Println(sec + "." + key)
	val, _ := i.file.String(sec + "." + key)
	return val
}
