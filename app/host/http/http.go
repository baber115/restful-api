package http

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"github.com/gin-gonic/gin"
)

var (
	handler = &Handler{}
)

func init() {
	app.RegistryGin(handler)
}

func (h *Handler) Config() {
	// 从IOC里面获取HostService的实例对象
	h.svc = app.GetImpl(host.AppName).(host.Service)
}

func (h *Handler) Name() string {
	return host.AppName
}

// 通过写一个实例类,把内部接口通过http协议暴露出去
// 需要内部接口的实现
// 该实例类会实现Gin和Handler
type Handler struct {
	svc host.Service
}

// 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
	r.DELETE("/hosts/:id", h.deleteHost)
	r.GET("/hosts", h.queryHost)
	r.GET("/hosts/:id", h.describeHost)
	r.PUT("/hosts/:id", h.putUpdateHost)
	r.PATCH("/hosts/:id", h.patchUpdateHost)
}
