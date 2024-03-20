package utils

import (
	"fmt"
	"wifi_plugins/common"
	"wifi_plugins/entity"
)

func LoadSubscribeLinkUtil(subscribeLinKPrefix string, cusInfo entity.CustomerInfo) (subscribeLink string) {
	//拼接SubscribeConfig的URL
	url := subscribeLinKPrefix + cusInfo.UserId + cusInfo.DeviceId + common.SubscribeConfigSuffix
	deviceConfig, err := ReadRemoteConfig(url)
	if err != nil {
		fmt.Printf("Read Remote Config Exception Message: %+v", err.Error())
	} else {
		link, err := FileContentDecrypt(deviceConfig)
		if err != nil {
			fmt.Printf("FileContentDecrypt Exception Message: %+v", err.Error())
		} else {
			subscribeLink = link
		}
	}
	return subscribeLink
}
