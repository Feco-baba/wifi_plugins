package utils

import (
	"encoding/json"
	Common "wifi_plugins/src/common"
	Entity "wifi_plugins/src/entity"
	ExcetionHandler "wifi_plugins/src/excetion"
)

func LoadDeviceConfigUtil(cusInfo Entity.CustomerInfo) (deviceBasicConfig Entity.DeviceBasicConfig, err error) {
	//拼接DeviceConfig的URL
	url := cusInfo.DeviceConfigLink + cusInfo.UserId + cusInfo.DeviceId + Common.DeviceSuffix
	deviceConfigJSONStr := ExcetionHandler.CommonExceptionHandler(ReadRemoteConfig(url))
	var config Entity.DeviceBasicConfig
	err = json.Unmarshal([]byte(deviceConfigJSONStr), &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
