package main

import (
	"fmt"

	"github.com/3Thiago/Dependency-Go/pkg/parte1"
	"github.com/3Thiago/Dependency-Go/pkg/parte2"
	"github.com/3Thiago/Dependency-Go/pkg/parte3"
)

func main() {
	fmt.Println("Hello World")
	teste := BuildParte()
	teste.GetParte1()
}

func BuildParte() *parte1.Parte1Struct {
	return &parte1.Parte1Struct{
		Parte2Interface: &parte2.Parte2Struct{},
		Parte3Interface: &parte3.Parte3Struct{},
	}
}
