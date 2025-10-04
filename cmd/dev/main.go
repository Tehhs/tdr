package main

import (
	"log"

	"github.com/Tehhs/tdr/pkg/core"
)

func main() { 
	//todo(code): should pick up on this
	coreInstance := core.NewTDRCore()
	processOutput, err := coreInstance.Process("./main.go")

	if err != nil { 
		panic(err)
	}

	for _, smth := range processOutput.Todos { 
		log.Printf("file %s", *smth.Name)
		
		for _, t := range smth.Processed { 
			log.Print(t)
		}
	}
}