package conf

// 全局config实例对象
// 也就是我们程序，在内存中的配置对象
// 程序内部获取配置，都通过读取该对象
// *配置加载时，config对象被初始化
//		LoadConfigFromToml
//		LoadConfigFromEnv
// 为了不被程序运行时恶意修改，设置成私有变量
var config *Config

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	Log   *Log   `toml:"log"`
	MySQL *MySQL `toml:"mysql"`
}

// 获取config配置的函数
func GetConfig() *Config {
	return config
}

// 初始化默认的config对象
func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		Log:   NewDefaultLog(),
		MySQL: NewDefaultMysql(),
	}
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
	// EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	// CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	// KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8005",
	}
}

// MySQL todo
type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// 因为使用的MySQL连接池，需要池做一些配置
	// 控制当前程序的MySQL打开的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制MySQL复用，比如5，最多运行5个复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 连接MySQL的生命周期，必须小于MySQL Server配置
	// 一个链接用12小时，必须换一个链接，保证一定的可用性
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// Idle链接，最多允许存活多久
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	// lock        sync.Mutex
}

func NewDefaultMysql() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "root",
		Password:    "root",
		Database:    "demo",
		MaxOpenConn: 10,
		MaxIdleConn: 5,
	}
}

// Log todo
// 用于配置全局对象
type Log struct {
	Level   string    `toml:"level" env:"LOG_LEVEL"`
	Format  LogFormat `toml:"format" env:"LOG_FORMAT"`
	To      LogTo     `toml:"to" env:"LOG_TO"`
	PathDir string    `toml:"path_dir" env:"LOG_PATH_DIR"`
}

func NewDefaultLog() *Log {
	return &Log{
		Level:  "info",
		Format: TextFormat,
		To:     ToStdout,
	}
}
