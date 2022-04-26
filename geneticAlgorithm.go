package main

import(
	"fmt"
	"math"
	"math/rand"
	"sort"
)


func operation(x, y, z float64)float64 {
	//return x * 9 - y * 4 + z / 3 - 25
	return x*3 - y*5 + z/2 - 5
}

func fitness(x, y, z float64)float64 {
	res := operation(x, y, z)

	if res == 0 {
		return 99999999.
	} else {
		return math.Abs(1/res)
	}
}


func main() {

	var solutions [1000][3]float64
	for solutionNum := 0; solutionNum <1000; solutionNum++ {
		value1 := rand.Float64() * 9999
		value2 := rand.Float64() * 9999
		value3 := rand.Float64() * 9999

		solutions[solutionNum] = [3]float64{value1, value2, value3}
	}

	for i := 0; i <1000000; i++ {
		var rankedSolutions [1000][4]float64
		var rankedValues []float64
		for sNum, s := range solutions {
			rankedSolutions[sNum] = [4]float64{fitness(s[0], s[1], s[2]), s[0], s[1], s[2]}
			rankedValues = append(rankedValues, fitness(s[0], s[1], s[2]))
		}

		sort.Float64s(rankedValues)

		var bestSolutions [100][4]float64
		for i, value := range rankedValues[900:] {
			for _, s := range rankedSolutions {
				if s[0] == value {
					bestSolutions[i] = s
				} else {
					continue
				}
			}
		}

                fmt.Printf("\n=-=-=-=-=-=-=-=-=-=-Best of generation %d-=-=-=-=-=-=-=-=-=\n", i)
                fmt.Println(bestSolutions[len(bestSolutions)-1])


		if bestSolutions[len(bestSolutions)-1][0] > 50000{
			break
		}

		var elements [100][3]float64
		for i, s := range bestSolutions {
			elements[i] = [3]float64{s[1], s[2], s[3]}
		}

		var newGen [1000][3]float64
		for i := 0; i <1000; i++ {
			e1 := elements[rand.Intn(100)][0] + rand.Float64() * 0.00001
			e2 := elements[rand.Intn(100)][1] + rand.Float64() * 0.00001
			e3 := elements[rand.Intn(100)][2] + rand.Float64() * 0.00001

			switch rand.Intn(6) {
			case 0:
				newGen[i] = [3]float64{-e1,e2,e3}

			case 1:
				newGen[i] = [3]float64{e1,-e2,e3}

			case 2:
				newGen[i] = [3]float64{e1,e2,-e3}

			case 3:
				newGen[i] = [3]float64{-e1,-e2,e3}

			case 4:
				newGen[i] = [3]float64{-e1,e2,-e3}

			case 5:
				newGen[i] = [3]float64{e1,-e2,-e3}

			case 6:
				newGen[i] = [3]float64{-e1,-e2,-e3}
			}
		}
		solutions = newGen
	}
}
