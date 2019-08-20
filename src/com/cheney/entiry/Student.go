package entiry

import "fmt"

type Student struct{
	People People
	no string
}

func (s Student) Remark()  {
	fmt.Print("student's work is study")
}