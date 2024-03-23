package operations

import (
	"log"
	"os/exec"
)

func TurnOnSSH(clientId string) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	log.Println("远程开启...")
	outCmd := exec.Command("./sunny", "clientid", clientId)
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		log.Println(err.Error())
	}
	outPutStr := string(output) + string("开启成功")
	log.Println(outPutStr)
}
