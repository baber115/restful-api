package conf_test

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/conf"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("demo", conf.GetConfig().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("APP_NAME", "demo2")
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		fmt.Println(conf.GetConfig().App.Name)
		should.Equal("demo2", conf.GetConfig().App.Name)
	}
}
