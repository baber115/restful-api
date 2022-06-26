package impl

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"context"
	"fmt"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/sqlbuilder"
)

// 业务处理层, controller层
func (i *HostServiceImpl) CreateHost(ctx context.Context, req *host.Host) (*host.Host, error) {
	//// 直接打印日志
	i.l.Named("Creaet").Debug("create host")
	i.l.Info("create host")
	//// 带format打印日志，等同于fmt.Sprintf()
	i.l.Debugf("create host %s", req.Name)
	//// 打印meta数据，常用于Trace系统
	i.l.With(logger.NewAny("request-id", "req01")).Debug("create host with meta kv")

	if err := i.save(ctx, req); err != nil {
		return req, err
	}
	return req, nil
}

func (i *HostServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.HostSet, error) {
	b := sqlbuilder.NewBuilder(queryHostSQL)
	if req.Keywords != "" {
		b.Where(
			"r.name like ? or r.description like ? or r.private_ip like ? or r.public_ip like ?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
			"1%",
			"2%",
		)
	}
	b.Limit(req.Offset(), req.GetPageNumber())
	querySql, args := b.Build()
	i.l.Debugf("query sql: %s, args: %v", querySql, args)

	// query stmt，构造一个prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySql)
	if err != nil {
		return nil, nil
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	set := host.NewHostSet()
	for rows.Next() {
		// 每扫描一行，就需要读取出来
		ins := host.NewHost()
		if err := rows.Scan(
			&ins.Id,
			&ins.Vendor,
			&ins.Region,
			&ins.CreateAt,
			&ins.ExpireAt,
			&ins.Type,
			&ins.Name,
			&ins.Description,
			&ins.Status,
			&ins.Tags,
			&ins.SyncAt,
			&ins.Account,
			&ins.PublicIP,
			&ins.PrivateIP,
			&ins.Resource,
			&ins.CPU,
			&ins.Memory,
			&ins.GPUAmount,
			&ins.GPUSpec,
			&ins.OSType,
			&ins.OSName,
			&ins.SerialNumber,
		); err != nil {
			return nil, err
		}
		set.Add(ins)
	}

	// total统计
	countSQL, args := b.BuildCount()
	i.l.Debugf("count sql: %s, args: %v", countSQL, args)
	countStmt, err := i.db.PrepareContext(ctx, countSQL)
	if err != nil {
		return nil, err
	}
	defer countStmt.Close()
	if err := countStmt.QueryRowContext(ctx, args...).Scan(&set.Total); err != nil {
		return nil, err
	}

	return set, nil
}

func (i *HostServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	b := sqlbuilder.NewBuilder(queryHostSQL)
	b.Where("r.id = ?", req.Id)

	querySql, args := b.Build()
	i.l.Debugf("describe sql: %s, args: %v", querySql, args)

	// query stmt，构造一个prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySql)
	if err != nil {
		return nil, nil
	}
	defer stmt.Close()

	ins := host.NewHost()
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
		&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
		&ins.Account, &ins.PublicIP, &ins.PrivateIP,
		&ins.ResourceId, &ins.CPU, &ins.Memory, &ins.GPUSpec, &ins.GPUAmount, &ins.OSType, &ins.OSName, &ins.SerialNumber,
	)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *HostServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	// 第一步，获取已有的对象
	ins, err := i.DescribeHost(ctx, host.NewDescribeHostRequestWithId(req.Id))
	if err != nil {
		return nil, err
	}

	// 第二步，判断更新模式，更新对象
	switch req.UpdateMode {
	case host.UPDATE_PUT_PUT:
		if err := ins.Put(req.Host); err != nil {
			return nil, err
		}
	case host.UPDATE_PUT_PATCH:
		if err := ins.Patch(req.Host); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("update_mode only required PUT/PATCH")
	}

	// 第三步，检查更新后的数据是否合法
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 第四步，更新数据库的数据
	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	// 第五步，返回更新后的对象
	return ins, nil
}

func (i *HostServiceImpl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	rsp, err := i.destroy(ctx, req)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
