package features

import (
	"fmt"
	"wifi_plugins/operations"
	"wifi_plugins/utils"
)

func RemoteSSHConnectionChecking() {
	//to instance customer information
	cusInfo, err := utils.LoadCustomerInfo()
	if err == nil {
		deviceBasicConfig, devConfigErr := utils.LoadDeviceConfigUtil(cusInfo)
		if devConfigErr == nil {
			fmt.Println("SshTurnOn:" + string(rune(deviceBasicConfig.SshTurnOn)))
			if deviceBasicConfig.SshTurnOn == 1 {
				operations.TurnOnSSH(cusInfo.ClientId)
			}
		} else {
			fmt.Printf("Get Device Config Exception Message: %+v", devConfigErr.Error())
		}
	} else {
		fmt.Printf("Get Customer Info Exception Message: %+v", err.Error())
	}
}
