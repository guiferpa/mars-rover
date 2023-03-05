package main

import (
	"bufio"
	"github/guiferpa/foxbit-challenge/domain/rover"
	"github/guiferpa/foxbit-challenge/handler/interface/cli"
	roversdk "github/guiferpa/foxbit-challenge/infra/sdk/rover"
	"os"
)

func main() {
	sdk := roversdk.NewRoverSDK()
	service := rover.NewUseCaseService(sdk)
	cli := cli.NewInterface(service)

	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	if err := cli.Run(input); err != nil {
		panic(err)
	}
}
