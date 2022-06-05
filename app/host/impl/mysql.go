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

func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host Service 服务的子Logger
		// 封装的Zap让其满足Logger接口
		// 为什么要封装：
		// 	1、Logger全局实例
		// 	2、Logger Level的动态调整，Logrus不支持Level共同调整
		// 	3、加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.GetConfig().MySQL.GetDB(),
	}
}

type HostServiceImpl struct {
	l  logger.Logger
	db *sql.DB
}

var impl = &HostServiceImpl{}

// _ import app 自动执行注册逻辑
func init() {
	app.Registry(impl)
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
