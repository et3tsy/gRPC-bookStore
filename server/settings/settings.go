package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 初始化 viper
func Init() (err error) {

	viper.SetConfigFile("./settings/config.yaml") // 配置文件路径

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		fmt.Printf("Fatal error config file: %v \n", err)
		return
	}

	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("settings have changed...\n")
	})
	return
}
