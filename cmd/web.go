package cmd

import (
	"log"
	"fmt"
	"net/http"

	"github.com/rakyll/statik/fs"
	"github.com/urfave/cli"

	_ "monomer-service-demo/statik"
)

func RunWeb(c *cli.Context) error {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("WEB-Service开始运行，通过*******进行访问！")

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	return http.ListenAndServe(":80", nil)
}
