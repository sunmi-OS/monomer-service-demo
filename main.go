package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
	"runtime"
	"monomer-service-demo/cmd"
	"github.com/sunmi-OS/gocore/sqlite"
)

const banner = `
                                                                     _                    _                                                            _ 
                                                                    (_)                  | |                                                          (_)
 ____   ___  ____   ___  ____  _____  ____     ___ _____  ____ _   _ _  ____ _____     __| |_____ ____   ___     _____ _____     ___ _   _ ____  ____  _ 
|    \ / _ \|  _ \ / _ \|    \| ___ |/ ___)   /___) ___ |/ ___) | | | |/ ___) ___ |   / _  | ___ |    \ / _ \   (_____|_____)   /___) | | |  _ \|    \| |
| | | | |_| | | | | |_| | | | | ____| |      |___ | ____| |    \ V /| ( (___| ____|  ( (_| | ____| | | | |_| |                 |___ | |_| | | | | | | | |
|_|_|_|\___/|_| |_|\___/|_|_|_|_____)_|      (___/|_____)_|     \_/ |_|\____)_____)   \____|_____)_|_|_|\___/                  (___/|____/|_| |_|_|_|_|_|

`
const (
	version = "1.0.0"
	name    = "monomer-service-demo"
)

func main() {

	// 打印log
	fmt.Print(banner)

	// 配置核心
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化数据库
	sqlite.NewDB("dbDefault")

	app := cli.NewApp()
	app.Name = name
	app.Email = "wenzhenxi@sunmi.com"
	app.Version = version

	app.Action = runAll

	app.Commands = []cli.Command{

		{
			Name:    "web",
			Aliases: []string{"w"},
			Usage:   "web静态业务服务",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开始运行WEB容器提供web访问",
					Action: cmd.RunWeb,
				},
			},
		}, {
			Name:    "api",
			Aliases: []string{"w"},
			Usage:   "api接口服务",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开发运行接口服务",
					Action: cmd.RunApi,
				},
			},
		}, {
			Name:    "db",
			Aliases: []string{"w"},
			Usage:   "数据服务",
			Subcommands: []cli.Command{
				{
					Name:   "CreateTable",
					Usage:  "run emq test qos0",
					Action: cmd.CreateTable,
				},
				{
					Name:   "UpdateTableV1_1",
					Usage:  "run emq test qos0",
					Action: cmd.UpdateTableV1_1,
				},
			},
		}, {
			Name:    "all",
			Aliases: []string{"w"},
			Usage:   "全部服务",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "所有服务同时运行",
					Action: runAll,
				},
			},
		}, {
			Name:    "benchmark",
			Aliases: []string{"w"},
			Usage:   "基准测试",
			Subcommands: []cli.Command{
				{
					Name:   "run_db_write",
					Usage:  "进行数据写入测试",
					Action: cmd.RunDBWrite,
				},
				{
					Name:   "run_db_read",
					Usage:  "进行数据读取入测试",
					Action: cmd.RunDBRead,
				},
			},
		},
	}

	app.Run(os.Args)

}

func runAll(c *cli.Context) error {
	cmd.CreateTable(c)
	go cmd.RunWeb(c)
	return cmd.RunApi(c)

}
