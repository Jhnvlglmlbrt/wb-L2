package main

import (
	"fmt"
	"strconv"
)

type StrategySort interface {
	Sort([]int)
}

type Context struct {
	strategy StrategySort
}

type BubbleSort struct {
}

// Sort sorts data.
func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

type InsertionSort struct {
}

// Sort sorts data.
func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}

func main() {

	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}

	ctx := new(Context)

	ctx.Algorithm(&BubbleSort{})
	ctx.Sort(data1)

	ctx.Algorithm(&InsertionSort{})
	ctx.Sort(data2)

	expect := "1,2,3,4,5,6,7,8,9,"

	var result1 string
	for _, val := range data1 {
		result1 += strconv.Itoa(val) + ","
	}

	if result1 != expect {
		fmt.Printf("Expect result1 to equal %s, but %s.\n", expect, result1)
	}
	fmt.Printf(" result1 equal %s, to %s.\n", expect, result1)

	var result2 string
	for _, val := range data2 {
		result2 += strconv.Itoa(val) + ","
	}

	if result2 != expect {
		fmt.Printf("Expect result2 to equal %s, but %s.\n", expect, result2)
	}
	fmt.Printf(" result2 equal %s, to %s.\n", expect, result1)
}
