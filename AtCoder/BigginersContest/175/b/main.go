package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	n := 0
	fmt.Fscanln(r, &n)
	arr := make([]int, 0)

	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Fscanf(r, "%d", &tmp)
		arr = append(arr, tmp)
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if len(arr) < 3 {
					return
				}
				// fmt.Println(len(arr))
				// fmt.Println(arr[i], arr[j], arr[k])
				// fmt.Println(n)
				// fmt.Println(i, j, k)

				if arr[i]+arr[j] > arr[k] && arr[i]+arr[k] > arr[j] && arr[j]+arr[k] > arr[i] && arr[i] != arr[j] && arr[i] != arr[k] && arr[j] != arr[k] {
					count++
				}
			}
		}
	}

	fmt.Fprintln(w, count)
}
