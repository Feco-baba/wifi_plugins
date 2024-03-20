package operations

import (
	"fmt"
	"os/exec"
)

func ResetServesConfig(folderName string, configURL string) (string, error) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	resetCMD := "s|Https='.*'|Https='" + configURL + "'|"
	pathCMD := "/data/" + folderName + "/configs/ShellCrash.cfg"
	outCmd := exec.Command("sed", "-i", resetCMD, pathCMD)
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		return "ResetServesConfig", fmt.Errorf(err.Error())
	}
	outPutStr := string(output) + string("设定配置成功")
	return outPutStr, nil
}
