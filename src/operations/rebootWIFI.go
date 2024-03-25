package operations

import (
	"log"
	"os/exec"
	"time"
)

func RebootWIFI() {
	time.Sleep(5 * time.Minute)
	outCmd := exec.Command("reboot")
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		log.Println("执行命令失败", err.Error())
	}
	log.Println(output)
}
