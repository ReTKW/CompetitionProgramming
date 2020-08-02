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
	k := 0
	fmt.Fscan(r, &k)
	count := 0

	mod := 0
	modMap := map[int]bool{}

	for {
		mod = (mod*10 + 7) % k
		count++

		if modMap[mod] {
			fmt.Fprintln(w, -1)
			return
		}

		if mod == 0 {
			fmt.Fprintln(w, count)
			return
		}
		modMap[mod] = true
	}
}
