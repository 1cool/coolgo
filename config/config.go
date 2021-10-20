package config

import (
	"github.com/spf13/viper"
	"log"
)

func SetUp(path string, filename string, fileType string) {
	//设置配置文件的名字
	viper.SetConfigName(filename)

	//viper.AddConfigPath("$HOME/.appname")  // 设置配置文件的搜索目录
	viper.AddConfigPath(path) // 设置配置文件和可执行二进制文件在用一个目录

	//设置配置文件类型
	viper.SetConfigType(fileType)

	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}

	viper.WatchConfig()
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
