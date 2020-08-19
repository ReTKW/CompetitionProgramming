package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	x := 0
	k := 0
	d := 0

	fmt.Fscanln(r, &x, &k, &d)
	ac := 0
	bc := 0
	a := 0
	b := 0

	c := x / d
	if c < k {
		k -= c
		x -= d * c
	}

	for i := 0; i < k; i++ {
		p := float64(x + d)
		m := float64(x - d)

		p = math.Abs(p)
		m = math.Abs(m)

		if m >= p {
			x = int(p)
			if a == x {
				ac++
			} else {
				ac = 0
				a = x
			}

		} else if p > m {
			x = int(m)
			if b == x {
				bc++
			} else {
				bc = 0
				b = x
			}
		}

		// fmt.Fprintln(w, a, b, ac, bc, x, m, p)
		if bc > 3 && ac > 3 {
			if (k-i)%2 == 1 {
				x = a
			} else if (k-i)%2 == 0 {
				x = b
			}
			break
		}
	}

	// fmt.Fprintln(w, x, k, d)
	fmt.Fprintln(w, x)
}
