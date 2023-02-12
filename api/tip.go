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

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	head = &ListNode{}
	tmp := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}

	// 遍历完成
	if list1 == nil {
		tmp.Next = list2
	} else {
		tmp.Next = list1
	}

	return head.Next
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists2(list1, list2.Next)
	return list2
}

// (0,0) (2,1)  (0,4)  (0,4)  (3,4)
// 0<2 DD 0<1 R !
// 2>0 UU 1<4 RRR !
// 0<3 DDD 0=0 !
// 0<3 DDD 4=4 !
// https://leetcode.cn/problems/alphabet-board-path/description/
func AlphabetBoardPath(target string) string {
	var stringAp string
	a, b := 0, 0
	for _, s := range target {
		x := int(s-97) / 5
		y := int(s-97) % 5
		m, n := a-x, b-y

		i := x
		j := y

		if m > 0 {
			for m > 0 {
				stringAp = stringAp + "U"
				m--
			}
		} else if m < 0 {
			if x == 5 {
				// 起始点不管在哪 向'z'移动, 都会进行最后一步D, 在此之前先移动到边界
				for m < -1 {
					stringAp = stringAp + "D"
					m++
				}
			} else {
				for m < 0 {
					stringAp = stringAp + "D"
					m++
				}
			}
		}

		if n > 0 {
			for n > 0 {
				stringAp = stringAp + "L"
				n--
			}
		} else if n < 0 {
			for n < 0 {
				stringAp = stringAp + "R"
				n++
			}
		}

		// 向'z'移动
		if x == 5 && m != 0 && n == 0 {
			stringAp = stringAp + "D"
		}

		stringAp = stringAp + "!"

		a = i
		b = j

	}

	return stringAp
}
