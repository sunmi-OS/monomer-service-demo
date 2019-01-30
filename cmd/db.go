package cmd

import (
	"github.com/urfave/cli"
	"github.com/sunmi-OS/gocore/sqlite"
	"fmt"
)

func CreateTable(c *cli.Context) error {

	fmt.Println("数据库初始化开始!")

	err := sqlite.GetORM().Exec(`
create table user (
    [id]            integer PRIMARY KEY autoincrement,
    [is_enable]     int default 1,
    [username]      varchar (50),
    [password]      varchar (50),
    [createdate]    datetime default (datetime('now', 'localtime')) 
);
`).Error

	fmt.Println("数据库初始化成功!")

	return err

}

func UpdateTableV1_1(c *cli.Context) error {

	fmt.Println("数据库更新开始!")

	err := sqlite.GetORM().Exec(`
ALTER TABLE user ADD COLUMN note Text;
`).Error

	fmt.Println("数据库更新完成!")

	return err
}
