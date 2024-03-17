package operations

import (
	"fmt"
	"os/exec"
)

func DeleteServerConfig() (string, error) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	outCmd := exec.Command("rm", "/data/clash/yamls/config.yaml")
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	outPutStr := string(output) + string("删除配置成功")
	return outPutStr, nil
}
