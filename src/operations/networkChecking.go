package operations

import (
	"net/http"
)

func CheckNetwork(url string) (string, error) {
	resp, err := http.Get("https://www.aliyun.com/")
	if err != nil {
		return "网络异常", err
	}
	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		return "网络请求失败：" + resp.Status, nil
	}
	defer resp.Body.Close()
	return "可以上网", nil
}
