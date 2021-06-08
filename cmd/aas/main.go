package main

import (
	"os"

	"github.com/patrickjahns/azure-account-switcher/pkg/command"
)

func main() {
	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
