package parte1

import (
	"fmt"

	parte2 "github.com/3Thiago/Dependency-Go/pkg/parte2"
	parte3 "github.com/3Thiago/Dependency-Go/pkg/parte3"
)

type Parte1Interface interface {
	GetParte1() string
}

type Parte1Struct struct {
	Parte2Interface parte2.Parte2Interface
	Parte3Interface parte3.Parte3Interface
}

func (p1 *Parte1Struct) GetParte1() string {
	fmt.Println("Getparte1")
	p1.Parte2Interface.GetParte2()
	p1.Parte3Interface.GetParte3()
	
	return "parte2" 
}
