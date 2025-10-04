package main

import (
	"log"

	"github.com/Tehhs/tdr/pkg/cli"
)

func main() { 
	log.Print("hi")

	cli.New(cli.NewCLIParams{})
}