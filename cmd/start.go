package cmd

import (
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app"
	"codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/conf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	// 引入所有服务的实例
	_ "codeup.aliyun.com/625e2dd5594c6cca64844075/restful-api-demo-07/app/all"
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

		app.InitImpl()

		// 提供一个Gin Router
		g := gin.Default()
		app.InitGin(g)
		return g.Run(conf.GetConfig().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f", "etc/demo.toml", "demo api 配置文件路径")
	RootCmd.AddCommand(StartCmd)
}
