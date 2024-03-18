package utils

import (
	"wifi_plugins/common"
	"wifi_plugins/entity"
	"wifi_plugins/excetion"
)

func LoadSubscribeLinkUtil(subscribeLinKPrefix string, cusInfo entity.CustomerInfo) (subscribeLink string) {
	//拼接SubscribeConfig的URL
	url := subscribeLinKPrefix + cusInfo.UserId + cusInfo.DeviceId + common.SubscribeConfigSuffix
	deviceConfig := excetion.CommonExceptionHandler(ReadRemoteConfig(url))
	link := excetion.CommonExceptionHandler(FileContentDecrypt(deviceConfig))
	return link
}
