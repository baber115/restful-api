package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// 加载全局实例
func LoadGlobal() (err error) {
	// 加载DB的全局实例
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return err
	}

	return nil
}

// 如果把配置映射成Config对象
// 从Toml的配置文件加载配置
func LoadConfigFromToml(filepath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filepath, config)
	if err != nil {
		return fmt.Errorf("load config from file error, path: %s,%s", filepath, err)
	}

	return nil
}

// 从环境变量加载配置
func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return nil
}
