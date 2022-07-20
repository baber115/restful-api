package conf_test

import (
	"fmt"
	"go-restful-api/conf"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		conf.GetConfig().MySQL.GetDB()
	}
}
