package features

import (
	"log"
	"wifi_plugins/common"
	"wifi_plugins/operations"
	"wifi_plugins/utils"
)

func RemoteSSHConnectionChecking() {
	//to instance customer information
	cusInfo, err := utils.LoadCustomerInfo()
	if err == nil {
		deviceBasicConfig, devConfigErr := utils.LoadDeviceConfigUtil(cusInfo)
		if devConfigErr == nil {
			if deviceBasicConfig.SshTurnOn == 1 && common.SSH_STATUS == 0 {
				log.Printf("deviceBasicConfig SshTurnOn %+v", deviceBasicConfig.SshTurnOn)
				log.Printf("common.SSH_STATUS %+v", common.SSH_STATUS)
				//标记SSH状态为1(已开启)
				common.SSH_STATUS = 1
				//开启SSH服务
				operations.TurnOnSSH(cusInfo.ClientId)
			}
			if deviceBasicConfig.SshTurnOn == 0 && common.SSH_STATUS == 1 {
				log.Printf("deviceBasicConfig SshTurnOn %+v", deviceBasicConfig.SshTurnOn)
				log.Printf("common.SSH_STATUS %+v", common.SSH_STATUS)
				//标记SSH状态为0(已关闭)
				common.SSH_STATUS = 0
			}
		} else {
			log.Printf("Get Device Config Exception Message: %+v", devConfigErr.Error())
		}
	} else {
		log.Printf("Get Customer Info Exception Message: %+v", err.Error())
	}
}
