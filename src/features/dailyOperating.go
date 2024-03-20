package features

import (
	"fmt"
	"wifi_plugins/operations"
	"wifi_plugins/utils"
)

func RefreshServerConfig() {
	//to instance customer information
	cusInfo, err := utils.LoadCustomerInfo()
	if err == nil {
		stopResult, stopErr := operations.StopClashService(cusInfo.FolderName)
		if err == nil {
			fmt.Printf(stopResult)
			deleteResult, deleteErr := operations.DeleteServerConfig(cusInfo.FolderName)
			if deleteErr == nil {
				fmt.Printf(deleteResult)
				deviceBasicConfig, devConfigErr := utils.LoadDeviceConfigUtil(cusInfo)
				fmt.Printf(deviceBasicConfig.ServerSubscribeURL)
				fmt.Printf(string(rune(deviceBasicConfig.SshTurnOn)))
				if devConfigErr == nil {
					subscribeLink := utils.LoadSubscribeLinkUtil(deviceBasicConfig.ServerSubscribeURL, cusInfo)
					fmt.Printf(subscribeLink)
					resetResult, resetErr := operations.ResetServesConfig(cusInfo.FolderName, subscribeLink)
					if resetErr == nil {
						fmt.Printf(resetResult)
						reloadResult, reloadErr := operations.StartClashService(cusInfo.FolderName)
						if reloadErr == nil {
							fmt.Printf(reloadResult)
							startResult, startErr := operations.StartClashService(cusInfo.FolderName)
							if startErr == nil {
								fmt.Printf(startResult)
							} else {
								fmt.Printf("Start Clash Exception Message: %+v", reloadErr.Error())
							}
						} else {
							fmt.Printf("Reload Exception Message: %+v", reloadErr.Error())
						}
					} else {
						fmt.Printf("Reset Exception Message: %+v", resetErr.Error())
					}
				} else {
					fmt.Printf("Get Device Config Exception Message: %+v", deleteErr.Error())
				}
			} else {
				fmt.Printf("Delete Exception Message: %+v", deleteErr.Error())
			}
		} else {
			fmt.Printf("Stop Clash Exception Message: %+v", stopErr.Error())
		}
	} else {
		fmt.Printf("Get Customer Info Exception Message: %+v", err.Error())
	}
}
