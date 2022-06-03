package http

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

// 用于暴露 host service 接口

func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	// 解析用户传来的参数
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
	}
	// 具体接口调用
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
	}
	response.Success(c.Writer, ins)
}

func (h *Handler) deleteHost(c *gin.Context) {
	req := &host.DeleteHostRequest{
		Id: c.Params.ByName("id"),
	}

	if err := c.Bind(req); err != nil {
		response.Failed(c.Writer, err)
	}
	resp, err := h.svc.DeleteHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
	}
	response.Success(c.Writer, resp)
}