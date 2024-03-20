package main

import (
	"fmt"
	"net/http"
	"time"
	http2 "wifi_plugins/http"
)

func main() {

	ticker := time.NewTicker(5 * time.Second)
	// 通过循环接收ticker的触发信号并执行任务
	go func() {
		for range ticker.C {
			fmt.Printf("1")
		}
	}()

	http.HandleFunc("/resetConfig", http2.ResetConfig) //重设节点配置
	http.HandleFunc("/turnOnSSH", http2.TurnOnSSH)     //重设节点配置
	err := http.ListenAndServe("0.0.0.0:8081", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}

	// 程序持续运行以保持ticker工作
	select {}

}

//func main() {
//启动定时任务
//c := cron.New()
//// 定义定时任务，每秒执行一次
//_, err := c.AddFunc("1/60 * * * * ?", func() { //每一分钟检查下clash是否过期
//	//TODO
//})
//if err != nil {
//	panic(err)
//}

// 开始执行定时任务
//c.Start()
