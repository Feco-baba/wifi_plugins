package entity

// 定义配置结构体
type CustomerInfo struct {
	UserId           string `json:"userId"`
	DeviceId         string `json:"deviceId"`
	DeviceConfigLink string `json:"deviceConfigLink"`
}
