package main

import (
	"fmt"
	"net/http"
	http2 "wifi_plugins/http"
)

func main() {
	http.HandleFunc("/resetConfig", http2.ResetConfig) //重设节点配置

	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
