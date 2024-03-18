package utils

import (
	"encoding/json"
	"wifi_plugins/common"
	"wifi_plugins/entity"
	"wifi_plugins/excetion"
)

func LoadDeviceConfigUtil(cusInfo entity.CustomerInfo) (deviceBasicConfig entity.DeviceBasicConfig, err error) {
	//拼接DeviceConfig的URL
	url := cusInfo.DeviceConfigLink + cusInfo.UserId + cusInfo.DeviceId + common.DeviceSuffix
	deviceConfigJSONStr := excetion.CommonExceptionHandler(ReadRemoteConfig(url))
	var config entity.DeviceBasicConfig
	err = json.Unmarshal([]byte(deviceConfigJSONStr), &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
