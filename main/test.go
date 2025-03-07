package main

/*
#cgo LDFLAGS: -L. -laverage
#cgo CFLAGS: -I.
#include <stdlib.h>

// 声明 C++ 计算平均数函数
extern double compute_average(int* data, int length);
*/
import "C"

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"runtime/trace"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
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

// 调用 C++ 计算平均数
func computeAverageWithCpp(data []int) float64 {
	if len(data) == 0 {
		return 0.0
	}

	// 确保 Go 侧数据转换为 C.int
	cData := make([]C.int, len(data))
	for i, v := range data {
		cData[i] = C.int(v) // 显式转换
	}

	// 传递 C 数组指针
	return float64(C.compute_average((*C.int)(unsafe.Pointer(&cData[0])), C.int(len(cData))))
}

// 生产者
func producer(buffer []int, ch chan<- int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range buffer {
		ch <- num // 发送数据
		// fmt.Println("Producer produces:", num)

		// 如果数据是 10，则阻塞，等待消费者计算
		if num == 10 {
			// fmt.Println("Producer pauses, waiting for consumer calculation...")
			<-pause // 生产者暂停
		}
		// time.Sleep(time.Second) // 模拟读取过程
	}
	close(ch) // 关闭通道，通知消费者没有新数据
}

// 生产者
func producer2(buffer []int, ch chan<- int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range buffer {
		ch <- num // 发送数据
		// fmt.Println("Producer2 produces:", num)

		// 如果数据是 10，则阻塞，等待消费者计算
		if num == 10 {
			// fmt.Println("Producer2 pauses, waiting for consumer calculation...")
			<-pause // 生产者暂停
		}
		// time.Sleep(time.Second) // 模拟读取过程
	}
	close(ch) // 关闭通道，通知消费者没有新数据
}

// 生产者
func producer3(buffer []int, ch chan<- int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range buffer {
		ch <- num // 发送数据
		// fmt.Println("Producer3 produces:", num)

		// 如果数据是 10，则阻塞，等待消费者计算
		if num == 10 {
			// fmt.Println("Producer3 pauses, waiting for consumer calculation...")
			<-pause // 生产者暂停
		}
		// time.Sleep(time.Second) // 模拟读取过程
	}
	close(ch) // 关闭通道，通知消费者没有新数据
}

// 生产者
func producer4(buffer []int, ch chan<- int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, num := range buffer {
		ch <- num // 发送数据
		// fmt.Println("Producer4 produces:", num)

		// 如果数据是 10，则阻塞，等待消费者计算
		if num == 10 {
			// fmt.Println("Producer4 pauses, waiting for consumer calculation...")
			<-pause // 生产者暂停
		}
		// time.Sleep(time.Second) // 模拟读取过程
	}
	close(ch) // 关闭通道，通知消费者没有新数据
}

// 消费者：计算 Go 版本的平均数
func consumer(ch <-chan int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var data []int
	for num := range ch {
		data = append(data, num)

		// 如果数据是 10，则计算平均数
		if num == 10 {
			// fmt.Println("Consumer 1 starts calculating the average...")
			avg := computeAverage(data)
			// time.Sleep(2 * time.Second) // 模拟计算过程
			fmt.Printf("Average 1 calculation result: %.2f\n", avg)
			// fmt.Println("Consumer 1 completes calculation and notifies the producer to continue...")

			// 发送信号，让生产者继续
			pause <- struct{}{}
		}
	}
}

// 消费者：计算 C++ 版本的平均数
func consumer2(ch <-chan int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var data []int
	for num := range ch {
		data = append(data, num)

		// 如果数据是 10，则计算平均数
		if num == 10 {
			// fmt.Println("Consumer 2 starts calculating the average using C++...")
			// time.Sleep(2 * time.Second) // 模拟计算过程
			avg := computeAverageWithCpp(data)
			fmt.Printf("Average 2 calculation result (C++): %.2f\n", avg)
			// fmt.Println("Consumer 2 completes calculation and notifies the producer to continue...")

			// 发送信号，让生产者继续
			pause <- struct{}{}
		}
	}
}

// 消费者：计算 C++ 版本的平均数
func consumer4(ch <-chan int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var data []int
	for num := range ch {
		data = append(data, num)

		// 如果数据是 10，则计算平均数
		if num == 10 {
			// fmt.Println("Consumer 4 starts calculating the average using C++...")
			// time.Sleep(2 * time.Second) // 模拟计算过程
			avg := computeAverageWithCpp(data)
			fmt.Printf("Average 4 calculation result (C++): %.2f\n", avg)
			// fmt.Println("Consumer 4 completes calculation and notifies the producer to continue...")

			// 发送信号，让生产者继续
			pause <- struct{}{}
		}
	}
}

// 消费者 3 (Python 计算平均数)
func consumer3(ch <-chan int, pause chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var data []int
	for num := range ch {
		data = append(data, num)

		if num == 10 {
			// fmt.Println("Consumer 3 starts calculating the average using Python...")
			// time.Sleep(2 * time.Second) // 模拟计算过程

			// 演示：通过启动 python 命令并通过命令行参数传递数据
			// 假设你有一个脚本 python_avg.py, 它从命令行参数里拿到一堆数字，然后打印结果。
			cmdArgs := make([]string, len(data)+1)
			cmdArgs[0] = "python_avg.py" // 你的 python 脚本的文件名
			for i, v := range data {
				cmdArgs[i+1] = strconv.Itoa(v)
			}

			// 启动 python 脚本
			cmd := exec.Command("python", cmdArgs...) // windows下也一样用"python"或"python.exe"
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Error executing Python script:", err)
				pause <- struct{}{}
				continue
			}
			// python 脚本把结果 print 到 stdout, 我们这里取出并转换成 float
			pyResult := strings.TrimSpace(string(output))
			avgValue, err := strconv.ParseFloat(pyResult, 64)
			if err != nil {
				fmt.Println("Failed to parse Python output:", err)
			} else {
				fmt.Printf("Average 3 calculation result (Python): %.2f\n", avgValue)
			}

			// fmt.Println("Consumer 3 completes calculation and notifies the producer to continue...")
			pause <- struct{}{}
		}
	}
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
	runtime.GOMAXPROCS(cpuNum - 2)
	// 创建 Trace 文件
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
	buffer := generateRandomArray(1000000, 1000)
	ch := make(chan int, 3)
	ch2 := make(chan int, 3)
	ch3 := make(chan int, 3) // 第三个通道，给 consumer3 用
	// ch4 := make(chan int, 3) // 第三个通道，给 consumer3 用

	pause := make(chan struct{})
	pause2 := make(chan struct{})
	pause3 := make(chan struct{})
	// pause4 := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(6)

	// 启动生产者
	go producer(buffer, ch, pause, &wg)
	go producer2(buffer, ch2, pause2, &wg)
	go producer3(buffer, ch3, pause3, &wg)
	// go producer4(buffer, ch4, pause4, &wg)

	// 启动消费者
	go consumer(ch, pause, &wg)
	go consumer2(ch2, pause2, &wg)
	go consumer3(ch3, pause3, &wg)
	// go consumer4(ch4, pause4, &wg)

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("All tasks completed")
}
