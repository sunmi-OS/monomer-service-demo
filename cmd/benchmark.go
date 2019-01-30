package cmd

import (
	"github.com/urfave/cli"
	"BITU-service/core/base"
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

		username := base.GetRandomString(10)
		password := base.GetRandomString(10)

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

		id := base.GetRandomNumeral(5)

		// 随机数据
		ModelUser.GetIdByInfo(cast.ToInt64(id))
		i++
	}

	return nil
}
