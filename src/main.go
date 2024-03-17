package main

import (
	"fmt"
	"net/http"
	WebServices "wifi_plugins/src/http"
)

func main() {
	http.HandleFunc("/resetConfig", WebServices.ResetConfig) //重设节点配置

	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
