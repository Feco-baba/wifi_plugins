package http

import (
	"fmt"
	"net/http"
	"os/exec"
	"wifi_plugins/features"
)

func Check(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://www.aliyun.com/")
	if err != nil {
		fmt.Fprintf(w, "无法访问外网: %s", err)
		return
	}
	defer resp.Body.Close()
	fmt.Fprintf(w, "可以上网")
}

func Restart(w http.ResponseWriter, r *http.Request) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	outCmd := exec.Command("sh", "/data/clash/start.sh", "start")
	output, err := outCmd.Output()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	outPutStr := string(output) + string("重启完了")
	fmt.Fprintf(w, outPutStr)
}

func Stop(w http.ResponseWriter, r *http.Request) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	outCmd := exec.Command("sh", "/data/clash/start.sh", "stop")
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	outPutStr := string(output) + string("Clash服务已停止")
	fmt.Fprintf(w, outPutStr)
}

func DeleteConfig(w http.ResponseWriter, r *http.Request) {
	//调用linux的命令
	// 创建一个新的 exec.Command 对象
	outCmd := exec.Command("rm", "/data/clash/yamls/config.yaml")
	// 获取命令输出
	output, err := outCmd.Output()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	outPutStr := string(output) + string("删除配置成功")
	fmt.Fprintf(w, outPutStr)
}

func ResetConfig(w http.ResponseWriter, r *http.Request) {
	features.RefreshServerConfig()
}

func TurnOnSSH(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("test")
	//features.RemoteSSHConnectionChecking()
}
