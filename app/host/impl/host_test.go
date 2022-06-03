package impl_test

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/conf"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host/impl"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	// 定义对象是满足该接口的实例
	service host.Service
)

func init() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}
	zap.DevelopmentSetup()
	// host service的具体实现
	service = impl.NewHostServiceImpl()
}

func TestCreate(T *testing.T) {
	should := assert.New(T)
	ins := host.NewHost()
	ins.Name = "test"
	ins, err := service.CreateHost(context.Background(), ins)

	if should.NoError(err) {
		fmt.Println(ins)
	}
}
