package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(FermatFactoring(200099)) // Adjust the large prime number to test different fermat factoring cases
}

func FermatFactoring(n float64) (float64, float64) {
	m := math.Sqrt(n)
	j := math.Ceil(m)
	fmt.Printf("j=%v m=%v\n", j, m)
	je := float64(j)
	i := 1000
	for je < n && i >= 0 {
		a := math.Pow(je, 2) - n
		fmt.Printf("a = %v\n", a)
		Floor_a := float64(math.Floor(a))
		if math.Sqrt(Floor_a) == math.Floor(math.Sqrt(Floor_a)) { // test and make sure that a is a whole number
			sqrta := math.Sqrt(Floor_a)
			fmt.Printf("sqrta = %v\n", sqrta)
			ret1 := je + sqrta
			ret2 := je - sqrta
			fmt.Printf("%v %v\n", ret1, ret2)
			return ret1, ret2
		}
		je++
		fmt.Printf("je = %v\n", je)
		i--
	}
	return 0, 0
}
