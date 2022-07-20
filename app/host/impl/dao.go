package impl

import (
	"context"
	"database/sql"
	"go-restful-api/app/host"
)

// 完成对象和数据之前的转换

func (i *HostServiceImpl) save(ctx context.Context, ins *host.Host) error {
	if err := ins.Validate(); err != nil {
		return err
	}
	var (
		err error
	)
	// 补充默认值
	ins.InjectDefault()

	// 用事务,插入host和resource表数据
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error,%s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error %s,", err)
			}
		}
	}()

	rstmt, err := tx.Prepare(InsertResourceSql)
	if err != nil {
		return err
	}
	// 如果这里不释放,会一直存留,如果多的话,再次使用会报错
	defer rstmt.Close()
	_, err = rstmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.CreateAt, ins.ExpireAt, ins.Type,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.Account, ins.PublicIP,
		ins.PrivateIP,
	)
	if err != nil {
		return err
	}

	dstmt, err := tx.Prepare(InsertDescribeSql)
	if err != nil {
		return err
	}
	// 如果这里不释放,会一直存留,如果多的话,再次使用会报错
	defer dstmt.Close()
	_, err = dstmt.Exec(
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec,
		ins.OSType, ins.OSName, ins.SerialNumber,
	)
	if err != nil {
		return err
	}

	return nil
}

func (i *HostServiceImpl) destroy(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	// 全局异常
	var (
		resStmt  *sql.Stmt
		descStmt *sql.Stmt
		err      error
	)
	//// 重新查询出来
	//ins, err := i.DesribeHost(ctx, host.NewDesribeHostRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error,%s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error %s,", err)
			}
		}
	}()

	descStmt, err = tx.Prepare(DeleteDescribeSql)
	if err != nil {
		return nil, err
	}
	defer descStmt.Close()
	_, err = descStmt.Exec(req.Id)
	if err != nil {
		return nil, err
	}

	resStmt, err = tx.Prepare(DeleteResourceSql)
	if err != nil {
		return nil, err
	}
	defer resStmt.Close()
	_, err = resStmt.Exec(req.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *HostServiceImpl) update(ctx context.Context, ins *host.Host) error {
	var err error
	// 开启事务
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if err := tx.Rollback(); err != nil {
				i.l.Error("rollback error,%s", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				i.l.Error("commit error %s,", err)
			}
		}
	}()

	// 更新resource表
	var (
		resStmt, hostStmt *sql.Stmt
	)
	resStmt, err = tx.PrepareContext(ctx, updateResourceSQL)
	if err != nil {
		return err
	}
	_, err = resStmt.Exec(ins.Vendor, ins.Region, ins.ExpireAt, ins.Name, ins.Description, ins.Id)
	if err != nil {
		return err
	}

	// 更新host表
	hostStmt, err = tx.PrepareContext(ctx, updateHostSQL)
	if err != nil {
		return err
	}
	_, err = hostStmt.Exec(ins.CPU, ins.Memory, ins.Id)
	if err != nil {
		return err
	}

	return nil
}
