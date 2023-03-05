package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github/guiferpa/mars-rover/domain/rover"
	"github/guiferpa/mars-rover/handler/interface/cli"
	roversdk "github/guiferpa/mars-rover/infra/sdk/rover"
)

var file string

func init() {
	flag.StringVar(&file, "file", "", "file input")
}

func buildInput() ([]string, error) {
	input := make([]string, 0)

	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			return []string{}, err
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			text := scanner.Text()
			input = append(input, text)
		}

		return input, nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	return input, nil
}

func main() {
	flag.Parse()

	sdk := roversdk.NewRoverSDK()
	service := rover.NewUseCaseService(sdk)
	cli := cli.NewInterface(service)

	input, err := buildInput()
	if err != nil {
		log.Println(err)
		return
	}

	if err := cli.Run(input); err != nil {
		log.Println(err)
	}
}
