package conf

import "fmt"

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
	// EnableSSL bool   `toml:"enable_ssl" env:"APP_ENABLE_SSL"`
	// CertFile  string `toml:"cert_file" env:"APP_CERT_FILE"`
	// KeyFile   string `toml:"key_file" env:"APP_KEY_FILE"`
}

func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func NewDefaultApp() *App {
	return &App{
		Name: "demo",
		Host: "127.0.0.1",
		Port: "8005",
	}
}
