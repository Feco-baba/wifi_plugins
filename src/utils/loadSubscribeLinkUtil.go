package utils

import (
	"log"
	"wifi_plugins/common"
	"wifi_plugins/entity"
)

func LoadSubscribeLinkUtil(subscribeLinKPrefix string, cusInfo entity.CustomerInfo) (subscribeLink string) {
	//拼接SubscribeConfig的URL
	url := subscribeLinKPrefix + cusInfo.UserId + cusInfo.DeviceId + common.SubscribeConfigSuffix
	deviceConfig, err := ReadRemoteConfig(url)
	if err != nil {
		log.Println("Read Remote Config Exception Message: ", err.Error())
	} else {
		link, err := FileContentDecrypt(deviceConfig)
		if err != nil {
			log.Println("FileContentDecrypt Exception Message: ", err.Error())
		} else {
			subscribeLink = link
		}
	}
	return subscribeLink
}
