package utils

import (
	"encoding/json"
	"wifi_plugins/common"
	"wifi_plugins/entity"
)

func LoadDeviceConfigUtil(cusInfo entity.CustomerInfo) (deviceBasicConfig entity.DeviceBasicConfig, err error) {
	//拼接DeviceConfig的URL
	url := cusInfo.DeviceConfigLink + cusInfo.UserId + cusInfo.DeviceId + common.DeviceSuffix
	deviceConfigJSONStr, err := ReadRemoteConfig(url)
	if err != nil {
		return deviceBasicConfig, err
	}
	var config entity.DeviceBasicConfig
	err = json.Unmarshal([]byte(deviceConfigJSONStr), &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
