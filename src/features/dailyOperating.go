package features

import (
	Entity "wifi_plugins/src/entity"
	ExceptionHandler "wifi_plugins/src/excetion"
	Operator "wifi_plugins/src/operations"
	Utils "wifi_plugins/src/utils"
)

func RefreshServerConfig() {
	//to instance customer information
	cusInfo := ExceptionHandler.CommonExceptionHandler(Utils.LoadCustomerInfo())
	// 停止Clash服务
	ExceptionHandler.CommonExceptionHandler(Operator.StopClashService())
	// 删除现有服务配置
	ExceptionHandler.CommonExceptionHandler(Operator.DeleteServerConfig())
	// 获取服务器订阅地址
	deviceBasicConfig := ExceptionHandler.CommonExceptionHandler[Entity.DeviceBasicConfig](Utils.LoadDeviceConfigUtil(cusInfo))
	subscribeLink := Utils.LoadSubscribeLinkUtil(deviceBasicConfig.ServerSubscribeURL, cusInfo)
	// 替换服务器订阅地址
	ExceptionHandler.CommonExceptionHandler(Operator.ResetServesConfig(subscribeLink))
	// 重新加载服务配置
	ExceptionHandler.CommonExceptionHandler(Operator.StartClashService())
	// 启动Clash服务
	ExceptionHandler.CommonExceptionHandler(Operator.StartClashService())
}
