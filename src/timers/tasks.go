package timers

import (
	"log"
	"time"
	"wifi_plugins/features"
	"wifi_plugins/operations"
	"wifi_plugins/utils"
)

func SSHConnectionCheckingTask() {
	for {
		// 等待5分钟后再次执行
		time.Sleep(5 * time.Minute)
		features.RemoteSSHConnectionChecking()
	}
}

func RefreshServerConfigAndRebootWIFITask() {
	cusInfo, err := utils.LoadCustomerInfo()
	if err != nil {
		log.Println("Get LoadCustomerInfo error:", err.Error())
	}
	// 计算每日reboot的时刻
	now := time.Now().UTC().Add(time.Hour * 8)
	dailyTime := time.Date(now.Year(), now.Month(), now.Day()+1, cusInfo.RebootTime, 0, 0, 0, time.UTC)
	// 计算等待时间
	waitDuration := dailyTime.Sub(now)
	log.Println("nowTime:", now.Format("2006-01-02 15:04:05"))
	log.Println("waitDuration:", waitDuration.Hours())
	time.AfterFunc(waitDuration, func() {
		features.RefreshServerConfig()
		operations.RebootWIFI()
	})
}
