// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	words := make(map[string]bool)
	for in.Scan() {
		w := strings.ToLower(in.Text())

		if words[w] {
			fmt.Println("TWICE!")
			return
		}
		words[in.Text()] = true
	}
}
