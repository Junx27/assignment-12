package main

import (
	"fmt"
	inttoroman "golang-day-12/intToRoman"
	romantoint "golang-day-12/romanToInt"
	smallestpositive "golang-day-12/smallestPositive"
	twosum "golang-day-12/twoSum"
)

func main() {
	fmt.Println(smallestpositive.SmallestPositive([]int{1, 2, 0}))         // Output: 3
	fmt.Println(smallestpositive.SmallestPositive([]int{3, 4, -1, 1}))     // Output: 2
	fmt.Println(smallestpositive.SmallestPositive([]int{7, 8, 9, 11, 12})) // Output: 1
	fmt.Println(smallestpositive.SmallestPositive([]int{1, 1, 0}))         // Output: 2

	fmt.Println(twosum.TwoSum([]int{2, 7, 11, 15}, 9)) // Output: [0, 1]
	fmt.Println(twosum.TwoSum([]int{3, 2, 4}, 6))      // Output: [1, 2]
	fmt.Println(twosum.TwoSum([]int{3, 3}, 6))         // Output: [0, 1]

	fmt.Println(romantoint.RomanToInt("III"))     // Output: 3
	fmt.Println(romantoint.RomanToInt("LVIII"))   // Output: 58
	fmt.Println(romantoint.RomanToInt("MCMXCIV")) // Output: 1994

	fmt.Println(inttoroman.IntToRoman(3749)) // Output: "MMMDCCXLIX"
	fmt.Println(inttoroman.IntToRoman(1994)) // Output: "MCMXCIV"
	fmt.Println(inttoroman.IntToRoman(58))   // Output: "LVIII"

}
