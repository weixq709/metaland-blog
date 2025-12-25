package config

import "github.com/spf13/viper"

var SysConfig = defaultConfig()

type SystemConfig struct {
	Port              int      `mapstructure:"port"`
	ContextPath       string   `mapstructure:"context_path"`
	ExcludeLoginPaths []string `mapstructure:"exclude_login_paths"`
}

func LoadConfig(configFile string) *viper.Viper {
	configLoader := viper.New()
	configLoader.SetConfigFile(configFile)
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
