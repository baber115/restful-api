package impl

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"context"
)

// 业务处理层, controller层
func (i *HostServiceImpl) CreateHost(ctx context.Context, req *host.Host) (*host.Host, error) {
	//// 直接打印日志
	//i.l.Named("Create").Debug("create host")
	//i.l.Debug("create host")
	//// 带format打印日志，等同于fmt.Sprintf()
	//i.l.Debugf("create host %s", req.Name)
	//// 打印meta数据，常用于Trace系统
	//i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")

	if err := i.save(ctx, req); err != nil {
		return req, err
	}
	return req, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.Describe) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req host.DeleteHostRequest) (*host.Host, error) {
	rsp, err := i.destroy(ctx, req)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
