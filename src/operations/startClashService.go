package operations

import (
	"fmt"
	"os/exec"
)

func StartClashService(folderName string) (string, error) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	path := "/data/" + folderName + "/start.sh"
	outCmd := exec.Command("sh", path, "start")
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		return "StartClashService", fmt.Errorf(err.Error())
	}
	outPutStr := string(output) + string("Clash服务已启用")
	return outPutStr, nil
}
