package app

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"fmt"
)

// IOC 容器层：管理所有的服务的实例

// 1.HostService的实例必须注册过来，HostService才有具体的实例
//		在服务启动的时候注册
// 2.Http 暴露模块，依赖Ioc里面的HostService
var (
	HostService host.Service
	implApps    = map[string]Service{}
)

func Registry(svc Service) {
	// 服务的实例注册到svcs map当中
	if _, ok := implApps[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registreid", svc.Name()))
	}
	implApps[svc.Name()] = svc
	// 更加对象满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}

// 用于初始化IOC，注册到IOC容器里的所有服务
func Init() {
	for _, v := range implApps {
		v.Config()
	}
}

type Service interface {
	Config()
	Name() string
}
