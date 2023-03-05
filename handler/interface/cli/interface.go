package cli

import (
	"fmt"
	"github/guiferpa/foxbit-challenge/domain/rover"
)

type Interface struct {
	usecase rover.UserCase
}

func (cli *Interface) Run(args []string) error {
	output, err := cli.usecase.Walk(3, 3, "E", "MMRMMRMRRM")
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}

func NewInterface(usecase rover.UserCase) *Interface {
	return &Interface{usecase}
}
