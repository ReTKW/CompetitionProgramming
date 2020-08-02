package main

import "fmt"

func main() {
	stones := make([]int, 0)
	n := 0
	q := 0
	fmt.Scanf("%d %d", &n, &q)
	c := 0
	l := 0
	r := 0

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &c)
		stones = append(stones, c)
	}

	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d", &l, &r)
		sli := stones[l:r]
		out := cC(sli)
		fmt.Println(out)

	}

}

func cC(sli []int) int {
	nums := make([]int, 0)
	count := 0
	only := true
	for _, s := range sli {

		for _, n := range nums {
			if n == s {
				only = false
				break
			}
		}

		if only {
			nums = append(nums, s)
			count++
		}
	}

	return count
}
