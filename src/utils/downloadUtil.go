package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ReadRemoteConfig(url string) (string, error) {
	var isNotSuccess bool = true
	var resp *http.Response
	var err error
	var req3time int = 3
	for isNotSuccess && req3time > 0 {
		time.Sleep(1 * time.Minute)
		resp, err = http.Get(url)
		if err != nil {
			log.Println("访问文件失败：", err.Error())
		} else if resp.StatusCode != 200 {
			log.Println("访问文件失败：", resp.StatusCode)
		} else {
			isNotSuccess = false
		}
		req3time--
	}
	defer resp.Body.Close()
	// 读取文件内容
	fileContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取文件内容失败：", err.Error())
		return "", err
	}
	return string(fileContent), nil
}
