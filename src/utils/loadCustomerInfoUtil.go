package utils

import (
	"github.com/spf13/viper"
	Entity "wifi_plugins/src/entity"
)

func LoadCustomerInfo() (cusInfo Entity.CustomerInfo, err error) {

	viper.SetConfigName("customerInfo") // 配置文件名称（不带扩展名）
	viper.AddConfigPath(".")            // 添加搜索路径（当前目录）

	// 读取配置文件
	err = viper.ReadInConfig() // 自动查找并读取名为config.yaml或config.yml的文件
	if err != nil {
		return cusInfo, err
	}

	// 将配置解析到结构体中
	err = viper.Unmarshal(&cusInfo)
	if err != nil {
		return cusInfo, err
	}

	// 输出配置信息
	return cusInfo, nil
}
