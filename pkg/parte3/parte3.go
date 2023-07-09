package parte3

import "fmt"

type Parte3Interface interface {
	GetParte3() string
	GetParte3FromParte1() string
	GetParte3FromParte2() string
}

type Parte3Struct struct {
}

func (p3 *Parte3Struct) GetParte3() string {
	fmt.Println("Parte3 in Getparte3")
	return "Parte3"
}

func (p *Parte3Struct) GetParte3FromParte1() string {
	fmt.Println("Parte3 in GetParte3FromParte1")
	return "GetParte3FromParte1"
}

func (p *Parte3Struct) GetParte3FromParte2() string {
	fmt.Println("Parte3 in GetParte3 From Parte2")
	return "GetParte3FromParte2"
}