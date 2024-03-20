package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ReadRemoteConfig(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("访问文件失败： %+v", err.Error())
	}
	defer resp.Body.Close()
	// 读取文件内容
	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取文件内容失败： %+v", err.Error())
	}
	return string(fileContent), nil
}
