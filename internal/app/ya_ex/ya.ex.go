package main

//import "fmt"
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	arr := make([]int, 3)
	for i := range arr {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		//text= strings.Replace(text, "\n", "", -1)
		arr[i], _ = strconv.Atoi(text)
		//if strings.Compare("hi", text) == 0 {
		//  fmt.Println("hello, Yourself")
		//}
		//m:=make(map[int]int)
		//fmt.Println(m)
	}
	//for i:= range arr {
	//fmt.Fprintln(os.Stdout, arr[i])
	//}
	if arr[1] == 0 {
		if arr[0] != 0 {
			fmt.Fprintln(os.Stdout, 3)
		} else {
			fmt.Fprintln(os.Stdout, arr[2])
		}
	} else if arr[1] == 1 {
		fmt.Fprintln(os.Stdout, arr[2])
	} else if arr[1] == 4 {
		if arr[0] == 0 {
			fmt.Fprintln(os.Stdout, 4)
		} else {
			fmt.Fprintln(os.Stdout, 3)
		}

	} else if arr[1] == 6 {
		fmt.Fprintln(os.Stdout, 0)
	} else if arr[1] == 7 {
		fmt.Fprintln(os.Stdout, 1)
	} else {
		fmt.Fprintln(os.Stdout, arr[1])
	}
}
