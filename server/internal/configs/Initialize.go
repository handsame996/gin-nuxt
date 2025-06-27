package configs

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitializeViper(c *ConfigModel) {
	configPath := getConfigPath()
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}

	if err = v.Unmarshal(c); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
		return
	}
}

func getConfigPath() (configPath string) {
	flag.StringVar(&configPath, "c", "", "config path")
	flag.Parse()
	if configPath != "" {
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", configPath)
		return
	}
	if env := os.Getenv("gin-nuxt-env"); env != "" {
		fmt.Printf("您正在使用环境变量 gin-nuxt-env 参数传递的值, config 的路径为 %s\n", configPath)
		return env
	}
	if configPath == "" {
		fmt.Printf("您正在使用默认路径 ./config.yaml 参数传递的值, config 的路径为 %s\n", configPath)
		return filepath.Clean("./config.yaml")
	}
	return
}
