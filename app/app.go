package app

import "codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"

// IOC 容器层：管理所有的服务的实例

// 1.HostService的实例必须注册过来，HostService才有具体的实例
//		在服务启动的时候注册
// 2.Http 暴露模块，依赖Ioc里面的HostService
var (
	HostService host.Service
)
