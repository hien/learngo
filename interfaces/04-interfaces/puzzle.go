// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type puzzle struct {
	title string
	price money
}

// if you remove this method,
// you can no longer add it to the `store` in `main()`.
func (p puzzle) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}
