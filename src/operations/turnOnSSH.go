package operations

import (
	"fmt"
	"os/exec"
)

func TurnOnSSH(clientId string) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	fmt.Printf("远程开启...")
	outCmd := exec.Command("./sunny", "clientid", clientId)
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	outPutStr := string(output) + string("开启成功")
	fmt.Printf(outPutStr)
}
