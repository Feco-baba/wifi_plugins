package features

import (
	"log"
	"wifi_plugins/entity"
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
			deviceBasicConfig, devConfigErr := utils.LoadDeviceConfigUtil(cusInfo)
			if devConfigErr == nil {
				if deviceBasicConfig.IsEnableServerConfig == 1 {
					deleteResult, deleteErr := operations.DeleteServerConfig(cusInfo.FolderName)
					if deleteErr == nil {
						log.Printf(deleteResult)
					} else {
						log.Printf("Delete Exception Message: %+v", deleteErr.Error())
						startClashService(cusInfo)
					}
					subscribeLink := utils.GetSubscribeLinkUtil(deviceBasicConfig.ServerSubscribeURL)
					log.Printf(subscribeLink)
					resetResult, resetErr := operations.ResetServesConfig(cusInfo.FolderName, subscribeLink)
					if resetErr == nil {
						log.Printf(resetResult)
						startClashService(cusInfo)
					} else {
						log.Printf("Reset Exception Message: %+v", resetErr.Error())
						startClashService(cusInfo)
					}
				} else {
					log.Printf("Servers is disable.")
				}
			} else {
				log.Printf("Get Device Config Exception Message: %+v", devConfigErr.Error())
			}
		} else {
			log.Printf("Stop Clash Exception Message: %+v", stopErr.Error())
		}
	} else {
		log.Printf("Get Customer Info Exception Message: %+v", err.Error())
	}
}

func startClashService(cusInfo entity.CustomerInfo) {
	reloadResult, reloadErr := operations.StartClashService(cusInfo.FolderName)
	if reloadErr == nil {
		log.Printf(reloadResult)
		startResult, startErr := operations.StartClashService(cusInfo.FolderName)
		if startErr == nil {
			log.Printf(startResult)
		} else {
			log.Printf("Start Clash Exception Message: %+v", startErr.Error())
		}
	} else {
		log.Printf("Reload Exception Message: %+v", reloadErr.Error())
	}
}
