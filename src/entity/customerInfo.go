package entity

// 定义配置结构体
type CustomerInfo struct {
	UserId           string `yaml:"userId"`
	DeviceId         string `yaml:"deviceId"`
	DeviceConfigLink string `yaml:"deviceConfigLink"`
}
