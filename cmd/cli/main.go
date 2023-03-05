package main

import (
	"bufio"
	"github/guiferpa/mars-rover/domain/rover"
	"github/guiferpa/mars-rover/handler/interface/cli"
	roversdk "github/guiferpa/mars-rover/infra/sdk/rover"
	"os"
)

func buildInput() []string {
	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	return input
}

func main() {
	sdk := roversdk.NewRoverSDK()
	service := rover.NewUseCaseService(sdk)
	cli := cli.NewInterface(service)

	input := buildInput()

	if err := cli.Run(input); err != nil {
		panic(err)
	}
}
