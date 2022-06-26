package impl_test

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host/impl"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/conf"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
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
	fmt.Println(zap.DevelopmentSetup())
	service = impl.NewHostServiceImpl()
}

func TestCreate(T *testing.T) {
	should := assert.New(T)
	ins := host.NewHost()
	ins.Name = "test"
	ins.Id = "ins-02"
	ins, err := service.CreateHost(context.Background(), ins)

	if should.NoError(err) {
		fmt.Println(ins)
	}
}

func TestDestroy(T *testing.T) {
	should := assert.New(T)
	req := host.DeleteHostRequest{}
	req.Id = "ins-02"
	ins, err := service.DeleteHost(context.Background(), &req)
	fmt.Println(ins, err)
	if should.NoError(err) {
		fmt.Println("destroy success")
	}
}

func TestQuery(T *testing.T) {
	should := assert.New(T)
	req := host.NewQueryHostRequest()
	req.Keywords = "接口测试"
	set, err := service.QueryHost(context.Background(), req)

	if should.NoError(err) {
		for i := range set.Items {
			fmt.Println(set.Items[i].Id)
		}
	}
}

func TestDescribe(T *testing.T) {
	should := assert.New(T)
	req := host.NewDescribeHostRequestWithId("ins-02")
	ins, err := service.DescribeHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}

func TestUpdate(T *testing.T) {
	should := assert.New(T)
	req := host.NewPatchUpdateHostRequest("ins-02")
	req.Name = "patch update"
	ins, err := service.UpdateHost(context.Background(), req)
	if should.NoError(err) {
		fmt.Println(ins.Id)
	}
}
