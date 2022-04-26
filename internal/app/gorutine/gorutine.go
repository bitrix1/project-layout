package main

import (
	"fmt"
	"sync"
	"time"
)
func makeTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}
func worker(id int) {
    // fmt.Printf("Worker %d starting\n", id)

    // time.Sleep(time.Microsecond * 100)
    // fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup
	t := makeTimestamp()
	fmt.Println("Start time: ", (makeTimestamp() - t)  )
    for i := 1; i <= 1000000; i++ {
        wg.Add(1)

        // i := i

        go func() {
            defer wg.Done()
			time.Sleep(time.Millisecond * 100)
            // worker(i)
        }()
    }

    wg.Wait()
	fmt.Println("End time: ", (makeTimestamp() - t)  )

}
// func f() {
// 	time.Sleep(100)
// 	// fmt.Println("+")
// }

// func main() {
// 	t := makeTimestamp()
// 	// fmt.Println("Start time: ", t)
// 	for i := 0; i < 10000; i++ {
// 		go f()
//     }
// 	time.Sleep(time.Second)
// 	fmt.Println("End time: ", (makeTimestamp() - t) / 1000 )
// }