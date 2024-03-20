package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"wifi_plugins/entity"
)

//func LoadCustomerInfo() (cusInfo entity.CustomerInfo, err error) {
//
//	viper.SetConfigName("customerInfo") // 配置文件名称（不带扩展名）github.com/spf13/viper windows 不能变易
//	viper.AddConfigPath(".")            // 添加搜索路径（当前目录）
//
//	// 读取配置文件
//	err = viper.ReadInConfig() // 自动查找并读取名为config.yaml或config.yml的文件
//	if err != nil {
//		return cusInfo, err
//	}
//
//	// 将配置解析到结构体中
//	err = viper.Unmarshal(&cusInfo)
//	if err != nil {
//		return cusInfo, err
//	}
//
//	// 输出配置信息
//	return cusInfo, nil
//}

func loadConfigFromJSON(file string) (*entity.CustomerInfo, error) {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("当前目录为：", dir)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var config entity.CustomerInfo
	if err = json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadCustomerInfo() (cusInfo entity.CustomerInfo, err error) {
	config, err := loadConfigFromJSON("customerInfo.json")
	if err != nil {
		log.Fatal(err)
	}

	// 使用配置
	log.Printf("LoadCustomerInfo:%+v", config)
	return *config, nil
}
