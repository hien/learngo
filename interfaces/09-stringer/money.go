// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import "fmt"

type money float64

// money satisfies the fmt.Stringer
func (m money) String() string {
	return fmt.Sprintf("$%.2f", m)
}
