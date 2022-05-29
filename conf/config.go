package conf

import "database/sql"

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

// 初始化默认的config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMysql(),
	}
}

// 全局config实例对象
// 也就是我们程序，在内存中的配置对象
// 程序内部获取配置，都通过读取该对象
// *配置加载时，config对象被初始化
//		LoadConfigFromToml
//		LoadConfigFromEnv
// 为了不被程序运行时恶意修改，设置成私有变量
var config *Config

// 获取config配置的函数
func GetConfig() *Config {
	return config
}

// 全局Mysql实例
var db *sql.DB

// 单例模式，一定要加锁，避免资源竞争
func (m *MySQL) GetDB() *sql.DB {
	// 直接加锁，销售临界区，这里是销售GetDB函数
	m.lock.Lock()
	defer m.lock.Unlock()

	if db == nil {
		// 如果DB不存在，生成新的实例
		conn, err := m.getDBConn()
		if err != nil {
			panic(err)
		}
		db = conn
	}

	return db
}
