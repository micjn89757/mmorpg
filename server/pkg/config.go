package pkg

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)

var (
	ProjectRootPath = path.Dir(getoncurrentPath()) + "/"
)

// 获取当前文件所在目录
func getoncurrentPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return path.Dir(filename)
}

// 创建Viper对象
func CreateConfig(file string) *viper.Viper {
	config := viper.New()
	configPath := ProjectRootPath + "common/conf/"
	// config.AddConfigPath(configPath) // 配置文件所在目录
	// config.SetConfigFile(file)       // 文件名
	// config.SetConfigType("toml")     // 文件类型
	configFile := configPath + file + ".toml"
	config.SetConfigFile(configFile)

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("找不到配置文件:%s", configFile)) // 系统初始化阶段发生任何错误，直接结束进程
		} else {
			panic(fmt.Errorf("解析配置文件%s出错:%s", configFile, err))
		}
	}

	return config
}
