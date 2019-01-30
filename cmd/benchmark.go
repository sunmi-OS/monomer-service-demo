package cmd

import (
	"github.com/urfave/cli"
	"github.com/sunmi-OS/gocore/utils"
	"fmt"
	"time"
	"github.com/spf13/cast"
)

func RunDBWrite(c *cli.Context) error {

	i := 0

	go func() {

		for {

			fmt.Println(i)

			time.Sleep(1 * time.Second)

		}
	}()

	for {

		username := utils.GetRandomString(10)
		password := utils.GetRandomString(10)

		// 随机数据
		ModelUser.CreateUser(username, password)
		i++
	}

	return nil
}

func RunDBRead(c *cli.Context) error {
	i := 0

	go func() {

		for {

			fmt.Println(i)

			time.Sleep(1 * time.Second)

		}
	}()

	for {

		id := utils.GetRandomNumeral(5)

		// 随机数据
		ModelUser.GetIdByInfo(cast.ToInt64(id))
		i++
	}

	return nil
}
