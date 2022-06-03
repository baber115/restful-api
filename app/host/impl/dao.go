package impl

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"context"
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
