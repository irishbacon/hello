package main

import (
	"fmt"

// 	"github.com/irishbacon/stringutil"
)

// func main() {
// 	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
// }

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() (ret int) {
		ret, a, b = a, b, a+b
		return
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// test


// package main

// import "golang.org/x/tour/pic"

// func Pic(dx, dy int) [][]uint8 {
//  var xs = make([]([]uint8), dy)
//     for y, _ := range xs {
//         xs[y] = make([]uint8, dx)
//         for x, _ := range xs[y] {
//             //xs[y][x] = uint8((x & y) & (x & y))
// 			xs[y][x] = uint8((x * y) & ( x * y))
//         }
//     }
    
//   return xs  

// }

// func main() {
//     pic.Show(Pic)
// }