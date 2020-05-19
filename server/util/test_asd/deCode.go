package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

var (
	list = []int{1, 3, 67, 8, 9, 95, 4, 35, 67, 8, 4, 3, 2, 2, 3, 4, 67, 5, 5, 67546, 234, 1}
)

func main() {

	var str = "wx:10365648"
	var b = []byte(str)
	var md5Str = sha1.Sum(b)
	fmt.Println(fmt.Sprintf("%x", md5Str))

	var w = sha256.New()
	io.WriteString(w, str)
	bw := w.Sum(nil)
	md5Str2 := hex.EncodeToString(bw)
	fmt.Println(md5Str2)

	fmt.Println(base64.StdEncoding.EncodeToString([]byte(str)))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(str)))

	fmt.Printf("冒泡排序:")
	fmt.Println(BubbleSort(list)) //冒泡排序
	fmt.Printf("选择排序:")
	fmt.Println(SelectSort(list)) //选择排序
	fmt.Printf("快速排序:")
	fmt.Println(QuickSort(list)) //快速排序
	fmt.Printf("插入排序:")
	fmt.Println(InsetSort(list)) //插入排序
	fmt.Printf("因为数组浅拷贝了，所以后边三个其实并没有重新排序\n")
	// fmt.Println(SleepSort(list))
	var factorialInt = 3
	fmt.Printf("factorial of %d is %d", factorialInt, factorial(factorialInt))
}

// 阶乘 小于n的所有正整数相乘的积 0的阶乘为1 n的阶乘写作n!, 阶乘用于表示n个数有多少种排列组合
func factorial(i int) int {
	if i < 1 {
		return 1
	}
	return factorial(i-1) * i
}

// SleepSort 睡眠排序 ，启动一个定时任务，以数组的key设定睡眠时间 （搞笑来的，不知道在实际项目中是否会有应用
func SleepSort(arr []int) []int {
	ch := make(chan int)
	for _, val := range arr {
		go func(val int) {
			time.Sleep(time.Duration(val) * 10000)
			ch <- val
		}(val)
	}
	for i := range arr {
		arr[i] = <-ch
	}
	return arr
}

// InsetSort 插入排序 假设一个有序数组 ，如果n小于有序数组的最大值，则将n插入，并且循环更新n的位置 直至插入n之后的新数组重新变成一个有序数组
func InsetSort(arr []int) []int {
	len := len(arr)
	for i := 1; i < len; i++ {
		var tmp = arr[i]
		for j := i - 1; j < i; j-- {
			if tmp < arr[j] {
				arr[j+1], arr[j] = arr[j], tmp
			} else {
				break
			}
		}
	}
	return arr
}

//BubbleSort 冒泡排序 通过两个循环 以lenght次遍历数组，比较arr[i] , arr[j]的值并且交换位置
func BubbleSort(arr []int) []int {
	len := len(arr)
	// 1-最后一个
	for i := 1; i < len; i++ {
		// 0 - 倒数第二个
		for j := 0; j < len-1; j++ {
			// 比较两组数据
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// SelectSort 选择排序 通过两个循环 查找最小的值 交换到最新到位置，类似与冒泡排序，但是更少的修改数据
func SelectSort(arr []int) []int {
	len := len(arr)
	// 1-倒数第二个
	for i := 0; i < len-1; i++ {
		k := i
		// 0 - 最后一个
		for j := i + 1; j < len; j++ {
			// 比较两组数据
			if arr[k] > arr[j] {
				k = j
			}
		}
		if k != i {
			arr[i], arr[k] = arr[k], arr[i]
		}
	}
	return arr
}

// QuickSort 选择一个基准数值，将目标分为左右两组，并且使用递归继续为左右两组排序，直至递归完成，拼接 left 基准数 right 返回新的数组
func QuickSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	baseNum := arr[0]
	left := []int{}
	right := []int{}

	for i := 1; i < len; i++ {
		if baseNum > arr[i] {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	// fmt.Printf("left:%v ,right: %v \n", left, right)
	left = QuickSort(left)
	right = QuickSort(right)

	left = append(left, baseNum)
	return append(left, right...)
}
