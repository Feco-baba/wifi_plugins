package main

import (
	"log"
	"net/http"
	"os"
	"wifi_plugins/common"
	http2 "wifi_plugins/http"
	"wifi_plugins/timers"
)

func init() {
	// 初始化SSH状态为0(已关闭)
	common.SSH_STATUS = 0
	file, err := os.Create("wifi_plugins.log")
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	log.SetOutput(file)
}

func main() {
	go timers.SSHConnectionCheckingTask()
	go timers.RefreshServerConfigAndRebootWIFITask()
	http.HandleFunc("/networkHealthCheck", http2.NetworkHealthCheck)
	http.HandleFunc("/resetConfig", http2.ResetConfig)
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		log.Println("http listen failed")
	}
	// 程序持续运行以保持ticker工作
	select {}

	//encryptStr, err := utils.FileContentEncrypt("https://l.db-link01.top/link/b7d9UfLB3BemmPIf?clash=1")
	//if err != nil {
	//}
	//fmt.Println(encryptStr)
	//decryptStr, err := utils.FileContentDecrypt(encryptStr)
	//if err != nil {
	//	fmt.Println("解密失败", err)
	//}
	//fmt.Printf(decryptStr)
}
