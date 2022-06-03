package cmd

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host/http"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/host/impl"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/conf"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	confFile string
	confType string
	confETCD string
)

// 程序启动时组装
// start command
// go run main.go start
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo 后端API",
	Long:  "启动 demo 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 加载程序配置
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			panic(err)
		}

		// 加载Host Service的实体类
		service := impl.NewHostServiceImpl()

		// 通过Hst API Handler 提供HTTP RestFul接口
		api := http.NewHostHTTPHandler(service)

		// 提供一个Gin Router
		g := gin.Default()
		api.Registry(g)
		g.Run(conf.GetConfig().App.HttpAddr())

		return errors.New("no flags find")
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
