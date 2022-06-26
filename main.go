package main

import (
	"fmt"
	"log"
	"time"
)

type I interface {
	Go() string
}

// bubble sort
func bubble(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// select sort nums using select sort algorithm
func selectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

func main() {
	mi := my{}
	fmt.Println(mi.Go())
	c := make(chan int)
	go func() {
		select {
		case _ = <-c:
			log.Println("goroutine1")
		}
	}()
	go func() {
		select {
		case _ = <-c:
			log.Println("goroutine2")
		}
	}()
	go func() {
		select {
		case _ = <-c:
			log.Println("goroutine3")
		}
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c <- 1
		log.Println("sent")
	}()
	time.Sleep(1 * time.Minute)
}

type my struct{}

func (m my) Go() string {
	return "wft"
}
