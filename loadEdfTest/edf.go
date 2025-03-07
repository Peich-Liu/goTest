package main

/*
#cgo LDFLAGS: -L. -ledfloader
#include <stdlib.h>

// 声明 C++ 函数
extern double* loadEdf(int* length);
extern void freeEdf(double* ptr);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var length C.int

	// 调用 C++ 的 loadEdf 函数
	dataPtr := C.loadEdf(&length)
	if dataPtr == nil {
		fmt.Println("Failed to load EDF data")
		return
	}
	defer C.freeEdf(dataPtr) // 确保内存释放

	// 将 C 数组转换为 Go 切片
	goData := make([]float64, length)
	for i := 0; i < int(length); i++ {
		goData[i] = float64(C.double(*(*C.double)(unsafe.Pointer(uintptr(unsafe.Pointer(dataPtr)) + uintptr(i)*unsafe.Sizeof(*dataPtr)))))
	}

	// 打印部分数据
	fmt.Println("Loaded EEG data:", goData)
}
