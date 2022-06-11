package conf

// LogFormat 日志格式
type LogFormat string

const (
	// TextFormat 文本格式
	TextFormat = LogFormat("text")
	// JSONFormat json格式
	JSONFormat = LogFormat("json")
)

// LogTo 日志记录到哪儿
type LogTo string

const (
	// ToFile 保存到文件
	ToFile = LogTo("file")
	// ToStdout 打印到标准输出
	ToStdout = LogTo("stdout")
)

// 用于配置全局对象
// LogTo 日志记录到哪儿 (ToFile,ToStdout)
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
