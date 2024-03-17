package utils

import (
	"io/ioutil"
	"net/http"
)

func ReadRemoteConfig(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// 读取文件内容
	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}
