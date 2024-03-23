package features

import (
	"log"
	"wifi_plugins/operations"
	"wifi_plugins/utils"
)

func RefreshServerConfig() {
	//to instance customer information
	cusInfo, err := utils.LoadCustomerInfo()
	if err == nil {
		stopResult, stopErr := operations.StopClashService(cusInfo.FolderName)
		if err == nil {
			log.Printf(stopResult)
			deleteResult, deleteErr := operations.DeleteServerConfig(cusInfo.FolderName)
			if deleteErr == nil {
				log.Printf(deleteResult)
				deviceBasicConfig, devConfigErr := utils.LoadDeviceConfigUtil(cusInfo)
				log.Printf(deviceBasicConfig.ServerSubscribeURL)
				log.Printf(string(deviceBasicConfig.SshTurnOn))
				if devConfigErr == nil {
					subscribeLink := utils.LoadSubscribeLinkUtil(deviceBasicConfig.ServerSubscribeURL, cusInfo)
					log.Printf(subscribeLink)
					resetResult, resetErr := operations.ResetServesConfig(cusInfo.FolderName, subscribeLink)
					if resetErr == nil {
						log.Printf(resetResult)
						reloadResult, reloadErr := operations.StartClashService(cusInfo.FolderName)
						if reloadErr == nil {
							log.Printf(reloadResult)
							startResult, startErr := operations.StartClashService(cusInfo.FolderName)
							if startErr == nil {
								log.Printf(startResult)
							} else {
								log.Printf("Start Clash Exception Message: %+v", reloadErr.Error())
							}
						} else {
							log.Printf("Reload Exception Message: %+v", reloadErr.Error())
						}
					} else {
						log.Printf("Reset Exception Message: %+v", resetErr.Error())
					}
				} else {
					log.Printf("Get Device Config Exception Message: %+v", deleteErr.Error())
				}
			} else {
				log.Printf("Delete Exception Message: %+v", deleteErr.Error())
			}
		} else {
			log.Printf("Stop Clash Exception Message: %+v", stopErr.Error())
		}
	} else {
		log.Printf("Get Customer Info Exception Message: %+v", err.Error())
	}
}
