package http

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"github.com/gin-gonic/gin"
)

var (
	API = &Handler{}
)

func NewHostHTTPHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Config() {
	if app.HostService == nil {
		panic("dependence host service is nil")
	}
	// 从IOC里面获取HostService的实例对象
	h.svc = app.HostService
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
}
