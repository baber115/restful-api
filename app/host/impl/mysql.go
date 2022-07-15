package impl

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/go-restful-api/app"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/go-restful-api/app/host"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/go-restful-api/conf"
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

// NewHostServiceImpl 保证调用该函数之前, 全局conf对象已经初始化
func NewHostServiceImpl() *HostServiceImpl {
	return &HostServiceImpl{
		// Host service 服务的子Loggger
		// 封装的Zap让其满足 Logger接口
		// 为什么要封装:
		// 		1. Logger全局实例
		// 		2. Logger Level的动态调整, Logrus不支持Level共同调整
		// 		3. 加入日志轮转功能的集合
		l:  zap.L().Named("Host"),
		db: conf.GetConfig().MySQL.GetDB(),
	}
}
