package main

import (
	"fmt"
	//"math"
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

/**
 *	Binary calculation
 */
func main() {
	var a int = 8
	var b int = 12
	fmt.Printf("Before Shifting: %d (%08b),%d (%08b)\n", (a), (a), (b), (b))
	fmt.Printf("After Shifting: %d (%08b),%d (%08b)", (a << 1), (a << 1), (b << 1), (b << 1))
	/*fmt.Println(needInt(Small))
	fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))*/
}

func addNums(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
>>>>>>> master
}

func swapBinary(x, y int) {

}
