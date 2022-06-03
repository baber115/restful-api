package impl

import (
	"context"

	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
)

// 业务处理层, controller层
func (i *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	//// 直接打印日志
	//i.l.Named("Create").Debug("create host")
	//i.l.Debug("create host")
	//// 带format打印日志，等同于fmt.Sprintf()
	//i.l.Debugf("create host %s", ins.Name)
	//// 打印meta数据，常用于Trace系统
	//i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")

	if err := ins.Validate(); err != nil {
		return nil, err
	}
	ins.InjectDefault()
	if err := i.save(ctx, ins); err != nil {
		return ins, err
	}

	return nil, nil
}

func (a *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	return nil, nil
}

func (a *HostServiceImpl) DescribeHost(ctx context.Context, req *host.Describe) (*host.Host, error) {
	return nil, nil
}

func (a *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (a *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
