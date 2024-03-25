package timers

import (
	"time"
	"wifi_plugins/features"
)

func SSHConnectionCheckingTask() {
	for {
		// 等待5分钟后再次执行
		time.Sleep(5 * time.Minute)
		features.RemoteSSHConnectionChecking()
	}
}

func RefreshServerConfigTask() {
	// 计算每日4点的时刻（假设是UTC时间）
	now := time.Now().UTC().Add(time.Hour * 8)
	dailyTime := time.Date(now.Year(), now.Month(), now.Day(), 4, 0, 0, 0, time.UTC)
	// 每次执行完后，重新计算下一天的4点
	for {
		// 计算等待时间
		waitDuration := dailyTime.Sub(now)
		if waitDuration.Hours() > 0 {
			time.AfterFunc(waitDuration, func() {
				features.RefreshServerConfig()
				// 更新下一次执行时间
				now = time.Now().UTC().Add(time.Hour * 8)
				dailyTime = time.Date(now.Year(), now.Month(), now.Day()+1, 4, 0, 0, 0, time.UTC)

			})
		}
		time.Sleep(15 * time.Minute)
	}
}
