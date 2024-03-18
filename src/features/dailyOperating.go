package excetionfeatures

import (
	"wifi_plugins/entity"
	"wifi_plugins/excetion"
	"wifi_plugins/operations"
	"wifi_plugins/utils"
)

func RefreshServerConfig() {
	//to instance customer information
	cusInfo := excetion.CommonExceptionHandler(utils.LoadCustomerInfo())
	// 停止Clash服务
	excetion.CommonExceptionHandler(operations.StopClashService())
	// 删除现有服务配置
	excetion.CommonExceptionHandler(operations.DeleteServerConfig())
	// 获取服务器订阅地址
	deviceBasicConfig := excetion.CommonExceptionHandler[entity.DeviceBasicConfig](utils.LoadDeviceConfigUtil(cusInfo))
	subscribeLink := utils.LoadSubscribeLinkUtil(deviceBasicConfig.ServerSubscribeURL, cusInfo)
	// 替换服务器订阅地址
	excetion.CommonExceptionHandler(operations.ResetServesConfig(subscribeLink))
	// 重新加载服务配置
	excetion.CommonExceptionHandler(operations.StartClashService())
	// 启动Clash服务
	excetion.CommonExceptionHandler(operations.StartClashService())
}
