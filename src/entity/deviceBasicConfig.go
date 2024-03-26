package entity

type DeviceBasicConfig struct {
	ServerSubscribeURL   string `json:"serverSubscribeURL"`
	SshTurnOn            int    `json:"sshTurnOn"`
	IsEnableServerConfig int    `json:"isEnableServerConfig"`
}
