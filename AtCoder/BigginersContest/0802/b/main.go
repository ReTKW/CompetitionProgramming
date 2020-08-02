package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	n := 0
	d := float64(0)

	count := 0
	x := float64(0)
	y := float64(0)

	fmt.Scanf("%d %f", &n, &d)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &x, &y)
		// x = rand.Float64() * 200000
		// y = rand.Float64() * 200000
		dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
		if dist <= float64(d) {

			count++
		}
	}

	fmt.Println(count)

}
