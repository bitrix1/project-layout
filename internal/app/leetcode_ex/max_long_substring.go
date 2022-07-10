package leetcode_ex

import "fmt"

func Test() {
	r := lengthOfLongestSubstring("aabcdbbbbd")
	fmt.Println("r=", r)
}

func lengthOfLongestSubstring(s string) int {

	result := -1

	dict := make(map[byte]int)
	cs := ""
	d := -1

	for i := 0; i < len(s); i++ {

		if v, ok := dict[s[i]]; !ok {
			cs += string(s[i]) //

			dict[s[i]] = i

			d++
			debug_ex("if !ok ", i, cs, d, v, s) //

		} else {
			cs += string(s[i]) //

			d = min(i-v-1, d+1)

			dict[s[i]] = i
			debug_ex("else ok true", i, cs, d, v, s) //
		}

		result = max(result, d)

	}
	fmt.Println(dict)
	return result + 1

}
func debug_ex(ok string, i int, cs string, d int, v int, s string) {
	fmt.Println("")
	fmt.Println("# ", ok, "############")
	fmt.Println("string=", s)
	fmt.Println("i=", i)
	fmt.Println("cs=", cs)
	fmt.Println("d=", d)
	fmt.Println("v=", v)
	fmt.Println("s[i]=", string(s[i]))
}

func max(a int, b int) int {

	if a > b {

		return a

	}

	return b

}

func min(a int, b int) int {

	if a < b {

		return a

	}

	return b

}
