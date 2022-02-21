package main

import "fmt"

func create_domain() []int {
	domain := make([]int, 0)
	for i := 1; i <= 100; i++ {
		domain = append(domain, i)
	}

	return domain
}

type Variable struct {
	Value  int
	Domain []int
	Empty  bool
}

func new_variable() Variable {
	return Variable{
		Domain: create_domain(),
		Empty:  true,
	}
}

type Constraint struct {
	Satisfied func(map[rune]Variable) bool
}

type Problem struct {
	Constraints []Constraint
}

func Problem1Constraints() []Constraint {
	C1 := Constraint{
		func(vars map[rune]Variable) bool {
			return vars['A'].Value == vars['B'].Value+vars['C'].Value+vars['D'].Value+vars['E'].Value+vars['F'].Value
		},
	}

	return []Constraint{C1}
}

func main() {
	fmt.Println("vim-go")

	vars := make(map[rune]Variable)
	for i := 'A'; i <= 'O'; i++ {
		vars[i] = new_variable()
		fmt.Println(vars[i])
	}

	//fmt.Println(vars['A'], vars['A'].Domain)
}
