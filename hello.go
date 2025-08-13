// package main

// import "fmt"

// func init() {
// 	fmt.Println("added new file")
// }

package main

import (
	"fmt"

	"rsc.io/quote"
)

func init() {
	fmt.Println(quote.Go())
}
