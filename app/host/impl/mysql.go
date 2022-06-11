package impl

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/conf"
	"database/sql"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	_ host.Service = &HostServiceImpl{}
)

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

var impl = &HostServiceImpl{}

// _ import app 自动执行注册逻辑
func init() {
	app.RegistryImpl(impl)
}

// 只要保证全局对象config和全局logger已经加载完成
func (i *HostServiceImpl) Config() {
	i.l = zap.L().Named("Host")
	i.db = conf.GetConfig().MySQL.GetDB()
}

// 返回服务名称
func (i *HostServiceImpl) Name() string {
	return host.AppName
}
