package api

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func fb() bool {
	return false
}

func TestSwitch() {
	switch fb() {
	case true:
		println("1")
	case false:
		println("0")
	default:
		println("-1")
	}
}

func Sum(values ...int) (sum int) {
	sum = 0
	for _, value := range values {
		sum += value
	}

	return
}

func TestFunctionMethods() {
	Sum()
	Sum(1)
	Sum(1, 2, 3)
	// 等价
	Sum([]int{}...) // <=> Sum(nil...)
	Sum([]int{1}...)
	Sum([]int{1, 2, 3}...)
}

// 交替打印数字和字母
func printChanNums() {
	num, letter := make(chan bool), make(chan bool)
	waitGroup := sync.WaitGroup{}
	go func() {
		i := 0
		for {
			select {
			case <-num:
				fmt.Print(i)
				letter <- true
				i++
				break
			default:
				break
			}
		}
	}()

	waitGroup.Add(1)

	go func(waitGroup *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					waitGroup.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++

				if i >= strings.Count(str, "") {
					i = 0
				}

				num <- true
				break
			default:
				break
			}
		}
	}(&waitGroup)

	num <- true
	waitGroup.Wait()
}

func Test_printChanNums() {
	printChanNums()
	fmt.Println()
}

// 三数之和
func threeSum(nums []int) (res [][]int) {
	lens := len(nums)
	if lens < 3 || nums == nil {
		return nil
	}
	sort.Ints(nums)

	for i, num := range nums {
		l, r := i+1, lens-1
		if num > 0 {
			return res
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l = l + 1
				}
				for l < r && nums[r-1] == nums[r] {
					r = r - 1
				}
				l += 1
				r -= 1
			} else if nums[i]+nums[l]+nums[r] < 0 {
				l += 1
			} else {
				r -= 1
			}

		}

	}

	return
}

func Test_threeNumberSum() {
	numberSum := threeSum([]int{-1, 0, 1, 2, -1, -4})

	for _, ints := range numberSum {
		for _, i := range ints {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
