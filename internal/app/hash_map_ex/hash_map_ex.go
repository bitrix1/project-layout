// https://stackoverflow.com/questions/28432398/difference-between-some-operators-golang
package hash_map_ex

import (
	"fmt"
	"unsafe"
)

// "hash/maphash"
// "runtime"

func test1() int {
	return 1
}
func createInt() *int {
	return new(int)
}
func createStr() *string {
	return new(string)
}
func Test() {
	fmt.Println("hash_map_ex")
	const size = unsafe.Sizeof(uint32(1))
	s1 := createStr()
	*s1 = "123123"
	s := s1 //createInt()
	// s := uintptr(s1) //createInt()

	fmt.Println(unsafe.Pointer(s))
	// fmt.Println("runtime.Func:", runtime.Func)
	test11 := 123
	fmt.Println(test11)
	// fmt.Println(test1())
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8
	// m["k2"] = 13

	// fmt.Println("map:", m)

	// v1 := m["k1"]
	// fmt.Println("v1: ", v1)

	// fmt.Println("len:", len(m))

	// delete(m, "k2")
	// fmt.Println("map:", m)

	// _, prs := m["k2"]
	// fmt.Println("prs:", prs)

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map:", n)

}

// Use bitwise OR | to get the bits that are in 1 OR 2
// 1     = 00000001
// 2     = 00000010
// 1 | 2 = 00000011 = 3
// fmt.Println(1 | 2)

// Use bitwise OR | to get the bits that are in 1 OR 5
// 1     = 00000001
// 5     = 00000101
// 1 | 5 = 00000101 = 5
// fmt.Println(1 | 5)

// Use bitwise XOR ^ to get the bits that are in 3 OR 6 BUT NOT BOTH
// 3     = 00000011
// 6     = 00000110
// 3 ^ 6 = 00000101 = 5
// fmt.Println(3 ^ 6)

// Use bitwise AND & to get the bits that are in 3 AND 6
// 3     = 00000011
// 6     = 00000110
// 3 & 6 = 00000010 = 2
// fmt.Println(3 & 6)

// Use bit clear AND NOT &^ to get the bits that are in 3 AND NOT 6 (order matters)
// 3      = 00000011
// 6      = 00000110
// 3 &^ 6 = 00000001 = 1
// fmt.Println(3 & 6)
