package main

import (
	"fmt"
	"sync"
)

// var (
// 	mu    sync.Mutex
// 	count int
// )

// func main() {
// 	// Start 1000 goroutines that increment the count variable
// 	for i := 0; i < 1000; i++ {
// 		go func() {
// 			mu.Lock()
// 			count++
// 			mu.Unlock()
// 		}()
// 	}

// 	// Wait for all goroutines to finish
// 	time.Sleep(1 * time.Second)

// 	// Print the final count
// 	fmt.Println("Final count:", count)
// }

func main() {
	// 互斥锁保护计数器
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	var wg sync.WaitGroup
	wg.Add(5)

	// 启动10个gourontine
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
