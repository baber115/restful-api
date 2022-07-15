package http

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/go-restful-api/app/host"
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

func (h *Handler) queryHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewQueryHostFromHTTP(c.Request)

	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (h *Handler) describeHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewDescribeHostRequestWithId(c.Param("id"))
	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.DescribeHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (h *Handler) putUpdateHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewPutUpdateHostRequest(c.Param("id"))
	// 解析body里的数据
	if err := c.Bind(req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")
	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}

func (h *Handler) patchUpdateHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewPatchUpdateHostRequest(c.Param("id"))
	// 解析body里的数据
	if err := c.Bind(req.Host); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	req.Id = c.Param("id")
	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.UpdateHost(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}
