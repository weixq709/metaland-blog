package config

import "github.com/spf13/viper"

var SysConfig = defaultConfig()

type SystemConfig struct {
	Port              int      `mapstructure:"port"`
	ContextPath       string   `mapstructure:"context_path"`
	ExcludeLoginPaths []string `mapstructure:"exclude_login_paths"`
}

func LoadConfig(configPath string) *viper.Viper {
	configLoader := viper.New()
	// 设置配置文件名称为config
	configLoader.SetConfigName("config")
	configLoader.SetConfigType("yaml")
	configLoader.AddConfigPath(configPath)
	if err := configLoader.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := configLoader.Unmarshal(&SysConfig); err != nil {
		panic(err)
	}
	return configLoader
}

func defaultConfig() *SystemConfig {
	return &SystemConfig{
		Port:              8080,
		ExcludeLoginPaths: []string{},
	}
}
