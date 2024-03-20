package operations

import (
	"fmt"
	"os/exec"
)

func StopClashService(folderName string) (string, error) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	path := "/data/" + folderName + "/start.sh"
	outCmd := exec.Command("sh", path, "stop")
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		return "StopClashService", fmt.Errorf(err.Error())
	}
	outPutStr := string(output) + string("Clash服务已停止")
	return outPutStr, nil
}
