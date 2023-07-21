package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	//"sort"
)

func main() {
	arr := []int{}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")

	s := strings.Split(text, " ")
	// s := strings.Split("3 3 2067", " ")
	for i := range s {
		e, _ := strconv.Atoi(s[i])
		arr = append(arr, e)
	}

	fmt.Println(arr)
	d := arr[0]
	m := arr[1]
	y := arr[2]
	y = y
	//   12 13 1970     13 12 2004
	//fmt.Print(d, m, y)
	if ((d <= 12 && m > 12) || (d > 12 && m <= 12) || (d == m && d <= 12)) && y >= 1970 && y <= 2069 {
		fmt.Println(1)
		// fmt.Print(n, i, j)
		//fmt.Print(a)
	} else {
		// fmt.Println("else")
		// fmt.Print(n, i, j)
		fmt.Print(0)

	}
}

func main_C() {
	// return errors.New("empty name")
	// y := math.Abs(math.Inf(-11))
	// fmt.Print(y)
	// x := int64(-2)
	// fmt.Println(i64.Abs(x))
	arr := []int{}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	fmt.Println(len(text))
	return
	// text = strings.Replace(string(text), "\r\n", "", -1)
	s := strings.Split(text, " ")
	for i := range s {
		e, _ := strconv.Atoi(s[i])
		arr = append(arr, e)
	}

	// fmt.Println(arr[0])
	sort.Ints(arr)
	// fmt.Println(arr[2])

	n := arr[2]
	i := arr[0]
	j := arr[1]
	fmt.Print(n, i, j)

	a := j - i - 1
	b := n - j + i - 1
	// if a < 0 {
	// 	a *= -1
	// }
	// if b < 0 {
	// 	b *= -1
	// }
	// a -= 1
	// b -= 1
	if (a) < (b) {
		// fmt.Println("if")
		// fmt.Print(n, i, j)
		fmt.Print(a)
	} else {
		// fmt.Println("else")
		// fmt.Print(n, i, j)
		fmt.Print(b)

	}
	// for i := range s {
	// }
	// for i:=0;i<3;i++ {
	// }
}
func main_A() {
	reader := bufio.NewReader(os.Stdin)
	arr := make([]int, 3)
	for i := range arr {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\r\n", "", -1)
		a, _ := strconv.Atoi(text)
		arr[i] = a
		//fmt.Print(arr, ok)
		//if strings.Compare("hi", text) == 0 {
		//  fmt.Println("hello, Yourself")
		//}
		//m:=make(map[int]int)
		//fmt.Println(m)
	}
	//for i:= range arr {
	//fmt.Print(arr[i])
	//}
	if arr[1] == 0 {
		if arr[0] != 0 {
			fmt.Print(3)
		} else {
			fmt.Print(arr[2])
		}
	} else if arr[1] == 1 {
		fmt.Print(arr[2])
	} else if arr[1] == 4 {
		if arr[0] == 0 {
			fmt.Print(4)
		} else {
			fmt.Print(3)
		}

	} else if arr[1] == 6 {
		fmt.Print(0)
	} else if arr[1] == 7 {
		fmt.Print(1)
	} else {
		fmt.Print(arr[1])
	}
}
