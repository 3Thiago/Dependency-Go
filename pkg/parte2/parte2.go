package parte2

import (
	"fmt"

	parte3 "github.com/3Thiago/Dependency-Go/pkg/parte3"
)

type Parte2Interface interface {
	GetParte2() string
}

type Parte2Struct struct {
	Parte3Interface parte3.Parte3Interface
}

func BuildParte() *Parte2Struct {
	return &Parte2Struct{
		Parte3Interface: &parte3.Parte3Struct{},
	}
}

func (p2 *Parte2Struct) GetParte2()  string{
	parts := BuildParte()
	fmt.Println("Getparte2 from parte1")
	parts.Parte3Interface.GetParte3FromParte2()
	return "\nteste"
}