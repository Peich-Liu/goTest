package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

// 计算 Go 端的平均数
func computeAverage(data []int) float64 {
	if len(data) == 0 {
		return 0.0
	}

	sum := 0
	for _, v := range data {
		sum += v
	}
	return float64(sum) / float64(len(data))
}

// 消费者：计算 Go 版本的平均数
func consumer(buffer []int, wg *sync.WaitGroup) {
	defer wg.Done()

	// fmt.Println("Consumer 1 starts calculating the average...")
	avg := computeAverage(buffer)
	// time.Sleep(2 * time.Second) // 模拟计算过程
	fmt.Printf("Average 1 calculation result: %.2f\n", avg)
	// fmt.Println("Consumer 1 completes calculation and notifies the producer to continue...")
}

func consumer4(buffer []int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println("Consumer 1 starts calculating the average...")
	avg := computeAverage(buffer)
	fmt.Printf("Average 1 calculation result: %.2f\n", avg)
	// fmt.Println("Consumer 1 completes calculation and notifies the producer to continue...")

}

func generateRandomArray(size int, maxVal int) []int {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间设置随机数种子，确保每次运行生成不同的随机数

	arr := make([]int, size) // 创建一个长度为 size 的数组
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxVal) // 生成 [0, maxVal) 之间的随机整数
	}
	return arr
}

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("当前机器的逻辑CPU数量为：", cpuNum)
	runtime.GOMAXPROCS(runtime.NumCPU()) // 创建 Trace 文件
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace file: %v", err)
	}
	defer f.Close()

	// 启用 Go Trace
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// buffer := []int{1, 3, 5, 10} // 生产者数据源
	buffer := generateRandomArray(1000000000, 1000)

	var wg sync.WaitGroup
	wg.Add(3)

	go consumer(buffer, &wg)
	go consumer(buffer, &wg)
	// go consumer2(ch2, pause2, &wg)
	// go consumer3(ch3, pause3, &wg)
	go consumer(buffer, &wg)

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("All tasks completed")
}
