package impl

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/conf"
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	_ host.Service = &HostServiceImpl{}
)

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		l:  zap.L().Named("Host"),
		db: conf.GetConfig().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}
