package main

import (
	"log"
	"oauth2-server/cmd"
	"os"

	"github.com/urfave/cli"
)

var (
	app        *cli.App
	configFile string
)

func init() {
	app = cli.NewApp()
	app.Name = "oauth2"
	app.Usage = "Oauth 2.0 Server"
	app.Author = "mwl"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "configFile",
			Value:       "config.yml",
			Destination: &configFile,
		},
	}
}
func main() {
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "创建数据库表",
			Action: func(c *cli.Context) error {
				return cmd.Init(configFile)
			},
		},
		{
			Name:  "loaddata",
			Usage: "加载demo数据",
			Action: func(c *cli.Context) error {
				return cmd.LoadData(c.Args(), configFile)
			},
		},
		{
			Name:  "run",
			Usage: "运行服务器",
			Action: func(c *cli.Context) error {
				return cmd.Run(configFile)
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
