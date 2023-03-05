package main

import (
	"github/guiferpa/foxbit-challenge/domain/rover"
	"github/guiferpa/foxbit-challenge/handler/interface/cli"
	roversdk "github/guiferpa/foxbit-challenge/infra/sdk/rover"

	"os"
)

func main() {
	sdk := roversdk.NewRoverSDK()
	service := rover.NewUseCaseService(sdk)
	cli := cli.NewInterface(service)

	args := os.Args[0:]
	if err := cli.Run(args); err != nil {
		panic(err)
	}
}
