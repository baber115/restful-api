package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 如果把配置映射成Config对象
// 从Toml的配置文件加载配置
func LoadConfigFromToml(filepath string) error {
	config = NewDefaultConfig()
	if _, err := toml.DecodeFile(filepath, config); err != nil {
		return fmt.Errorf("load config from file error, path: %s,%s", filepath, err)
	}

	return nil
}

// 从环境变量加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	if err := env.Parse(config); err != nil {
		return err
	}
	return nil
}
