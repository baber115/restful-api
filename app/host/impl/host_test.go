package impl_test

import (
	"context"
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
	zap.DevelopmentSetup()
	// host service的具体实现
	service = impl.NewHostServiceImpl()
}

func TestCreate(T *testing.T) {
	ins := host.NewHost()
	ins.Name = "test"
	service.CreateHost(context.Background(), ins)
}
