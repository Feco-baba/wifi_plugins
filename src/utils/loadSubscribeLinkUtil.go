package utils

import (
	Common "wifi_plugins/src/common"
	Entity "wifi_plugins/src/entity"
	ExcetionHandler "wifi_plugins/src/excetion"
)

func LoadSubscribeLinkUtil(subscribeLinKPrefix string, cusInfo Entity.CustomerInfo) (subscribeLink string) {
	//拼接SubscribeConfig的URL
	url := subscribeLinKPrefix + cusInfo.UserId + cusInfo.DeviceId + Common.SubscribeConfigSuffix
	deviceConfig := ExcetionHandler.CommonExceptionHandler(ReadRemoteConfig(url))
	link := ExcetionHandler.CommonExceptionHandler(FileContentDecrypt(deviceConfig))
	return link
}
