package app

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"fmt"
	"github.com/gin-gonic/gin"
)

// IOC 容器层：管理所有的服务的实例

// 1.HostService的实例必须注册过来，HostService才有具体的实例
//		在服务启动的时候注册
// 2.Http 暴露模块，依赖Ioc里面的HostService
var (
	HostService host.Service
	implApp     = map[string]ImplService{}
	ginApp      = map[string]GinService{}
)

func GetImpl(name string) interface{} {
	for k, v := range implApp {
		if k == name {
			return v
		}
	}

	return nil
}

func RegistryImpl(svc ImplService) {
	// 服务的实例注册到svcs map当中
	if _, ok := implApp[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registreid", svc.Name()))
	}
	implApp[svc.Name()] = svc
	// 更加对象满足的接口来注册具体的服务
	if v, ok := svc.(host.Service); ok {
		HostService = v
	}
}
func RegistryGin(svc GinService) {
	// 服务的实例注册到svcs map当中
	if _, ok := ginApp[svc.Name()]; ok {
		panic(fmt.Sprintf("service %s has registreid", svc.Name()))
	}
	ginApp[svc.Name()] = svc
}

// 用于初始化IOC，注册到IOC容器里的所有服务
func InitImpl() {
	for _, v := range implApp {
		v.Config()
	}
}

// 用于初始化IOC，注册到IOC容器里的所有服务
func InitGin(r gin.IRouter) {
	// 先初始化所有对象
	for _, v := range ginApp {
		v.Config()
	}

	// 再完成http handler的注册
	for _, v := range ginApp {
		v.Registry(r)
	}
}

type ImplService interface {
	Config()
	Name() string
}

// 注册gin编写的handler
type GinService interface {
	Registry(r gin.IRouter)
	Config()
	Name() string
}
