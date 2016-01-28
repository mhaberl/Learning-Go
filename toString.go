package main

import "fmt"

type Matrix2x2 struct {
	a00, a01, a10, a11 int
}

func (m Matrix2x2) String() string {
	return fmt.Sprintf("%v\t%v\n%v\t%v\n", m.a00, m.a01, m.a10, m.a11)
}

func main() {

	m:=Matrix2x2{2,1,1,4}
	fmt.Println(m)
}
