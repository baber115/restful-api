package host

import "context"

// host app service 接口定义
type Service interface {
	// 创建
	CreateHost(context.Context, *Host) (*Host, error)
	// 列表查询
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	// 详情查询
	DescribeHost(context.Context, *Describe) (*Host, error)
	// 更新
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	// 删除
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}
