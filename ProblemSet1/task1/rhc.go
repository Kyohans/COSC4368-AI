package main

import (
	"fmt"
	. "math"
	"math/rand"
)

func function(x float64, y float64) float64 {
	return (-1)*(y+47)*Sin(Sqrt(Abs((x/2)+(y+47)))) - x*Sin(Sqrt(Abs(x-(y+47))))
}

func RHC(sp_x float64, sp_y float64, p int, z float64, seed int64) [][]float64 {
	sol_list := make([][]float64, 0)
	rand.Seed(seed)

	max, min := 512.0, -512.0
	for i := 1; i < p; i++ {
		z1 := (-1)*z + rand.Float64()*(z-(-1)*z)
		z2 := (-1)*z + rand.Float64()*(z-(-1)*z)

		for sp_x+z1 <= min && sp_x+z1 >= max && sp_y+z2 <= min && sp_y+z2 >= max {
			z1 = (-1)*z + rand.Float64()*(z-(-1)*z)
			z2 = (-1)*z + rand.Float64()*(z-(-1)*z)
		}

		sol_list = append(sol_list, []float64{function(sp_x+z1, sp_y+z2), sp_x + z1, sp_y + z2})
	}

	return sol_list
}

func minimize_func(sol_list [][]float64) []float64 {
	local_min := sol_list[0]

	for _, s := range sol_list {
		if local_min[0] > s[0] {
			local_min = s
		}
	}

	return local_min
}

func evaluate_rhc(sp_x float64, sp_y float64, p int, z float64, seed int64, count int) (int, []float64, float64) {
	count += 1

	sol_list := RHC(sp_x, sp_y, p, z, seed)
	local_min := minimize_func(sol_list)

	new_sols := RHC(local_min[1], local_min[2], p, z, seed)
	new_min := minimize_func(new_sols)

	if local_min[0] < new_min[0] {
		return count, []float64{local_min[1], local_min[2]}, local_min[0]
	} else {
		return evaluate_rhc(local_min[1], local_min[2], p, z, seed, count)
	}
}

func main() {
	sp := [][]float64{
		{404.0, 504.0},
		{0.0, 0.23},
		{-200.0, 300.0},
		{412.0, -99.9},
	}

	for _, i := range sp {
		fmt.Printf("f(%v, %v): %v\n", i[0], i[1], function(i[0], i[1]))
	}

	p := []int{30, 250}
	z := []float64{3.0, 0.5}

	var num_runs int
	for seedi := 0; seedi < 2; seedi++ {
		randseed := rand.Int63()
		fmt.Printf("\nRun %d with random seed %d\n\n", seedi+1, randseed)

		for _, v := range sp {
			x, y := v[0], v[1]

			for _, neighbor_count := range p {
				for _, size := range z {

					fmt.Printf("Run %d for parameters sp: %v, p: %d, z: %v\n", num_runs+1, v, neighbor_count, size)

					count, vector, solution := evaluate_rhc(x, y, neighbor_count, size, randseed, 0)
					num_runs += 1

					fmt.Printf("\tCount: %d, Sol: %v, f(sol): %v\n", count, vector, solution)
				}
			}
		}
	}

	fmt.Printf("\n33rd run with parameters: \n\tsp: {267.15, 510.125}\n\tp: 45\n\tz: 0.3")
	hcount, hvector, hsolution := evaluate_rhc(267.15, 510.125, 45, 0.3, rand.Int63(), 0)

	fmt.Printf("\nResults: \n\tCount: %d\n\tSol: %v\n\tf(sol): %v\n", hcount, hvector, hsolution)
}
