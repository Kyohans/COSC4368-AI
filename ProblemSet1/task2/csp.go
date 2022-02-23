package main

import (
	"fmt"
	"math"

	"github.com/gnboorse/centipede"
)

func main() {
	vars := centipede.Variables{
		centipede.NewVariable("A", centipede.IntRange(1, 101)),
		centipede.NewVariable("B", centipede.IntRange(1, 101)),
		centipede.NewVariable("C", centipede.IntRange(1, 101)),
		centipede.NewVariable("D", centipede.IntRange(1, 101)),
		centipede.NewVariable("E", centipede.IntRange(1, 101)),
		centipede.NewVariable("F", centipede.IntRange(1, 101)),
		centipede.NewVariable("G", centipede.IntRange(1, 101)),
		centipede.NewVariable("H", centipede.IntRange(1, 101)),
		centipede.NewVariable("I", centipede.IntRange(1, 101)),
		centipede.NewVariable("J", centipede.IntRange(1, 101)),
		centipede.NewVariable("K", centipede.IntRange(1, 101)),
		centipede.NewVariable("L", centipede.IntRange(1, 101)),
		centipede.NewVariable("M", centipede.IntRange(1, 101)),
		centipede.NewVariable("N", centipede.IntRange(1, 101)),
		centipede.NewVariable("O", centipede.IntRange(1, 101)),
	}

	constraints := centipede.Constraints{
		centipede.Constraint{Vars: centipede.VariableNames{"A", "B", "C", "D", "E", "F"},
			ConstraintFunction: func(variables *centipede.Variables) bool {
				if variables.Find("A").Empty || variables.Find("B").Empty || variables.Find("C").Empty || variables.Find("D").Empty || variables.Find("E").Empty || variables.Find("F").Empty {
					return true
				}
				return variables.Find("A").Value.(int) == variables.Find("B").Value.(int)+variables.Find("C").Value.(int)+variables.Find("D").Value.(int)+variables.Find("E").Value.(int)+variables.Find("F").Value.(int)
			}},
		centipede.Constraint{Vars: centipede.VariableNames{"A", "D", "E"},
			ConstraintFunction: func(variables *centipede.Variables) bool {
				if variables.Find("A").Empty || variables.Find("D").Empty || variables.Find("E").Empty {
					return true
				}
				return variables.Find("D").Value.(int) == int(math.Sqrt(float64(variables.Find("E").Value.(int)*variables.Find("E").Value.(int)*variables.Find("A").Value.(int)+694)))
			}},
	}

	solver := centipede.NewBackTrackingCSPSolver(vars, constraints)
	solver.State.MakeArcConsistent()
	success := solver.Solve()
	fmt.Println(success)
}
