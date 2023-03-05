package cli

import (
	"fmt"
	"github/guiferpa/foxbit-challenge/domain/rover"
	"strconv"
	"strings"
)

type Interface struct {
	usecase rover.UserCase
}

func (cli *Interface) Run(input []string) error {
	slot := input[1:]

	for i := 0; i < len(slot); i += 2 {
		position := strings.Fields(slot[i])

		x, err := strconv.Atoi(position[0])
		if err != nil {
			return err
		}

		y, err := strconv.Atoi(position[1])
		if err != nil {
			return err
		}

		dir := position[2]

		inst := slot[i+1]

		output, err := cli.usecase.Walk(x, y, dir, inst)
		if err != nil {
			return err
		}

		fmt.Println(output)
	}

	return nil
}

func NewInterface(usecase rover.UserCase) *Interface {
	return &Interface{usecase}
}
