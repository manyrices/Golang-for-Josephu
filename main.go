package main

import (
	"fmt"
)

type Boy struct {
	no   int
	next *Boy
}

func main() {
	var total, star, count int
	fmt.Printf("How many child play game: ")
	fmt.Scanf("%d\n", &total)
	//showQueue(head)
	for {
		fmt.Printf("Few child start: ")
		fmt.Scanf("%d\n", &star)
		if star < total {
			break
		} else {
			fmt.Printf("star number %d bigger than total number %d..\n", star, total)
		}
	}
	head := initQueue(total)
	fmt.Printf("set a number to play: ")
	fmt.Scanf("%d\n", &count)
	fmt.Println("\nGame start........")
	Game(head, star, count)
	fmt.Scanf("")
}

func initQueue(num int) *Boy {
	head := &Boy{}
	temp := head
	if num < 2 {
		fmt.Println("Too few people to complete the game.")
		return head
	}
	for i := 1; i <= num; i++ {
		curBoy := &Boy{
			no:   i,
			next: head,
		}
		if i == 1 {
			head = curBoy
		}
		temp.next = curBoy
		temp = temp.next
	}
	return head
}

func showQueue(head *Boy) {
	if head.next == nil {
		fmt.Println("Queue is empty...")
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		fmt.Printf("Boy:%d => ", temp.no)
		temp = temp.next
	}
	fmt.Println()
}

func Game(head *Boy, startNum int, count int) {
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	for i := 1; i <= startNum-1; i++ {
		head = head.next
		temp = temp.next
	}
	for {
		for i := 1; i <= count-1; i++ {
			head = head.next
			temp = temp.next
		}
		fmt.Printf("Boy:%d out\n", head.no)
		head = head.next
		temp.next = head

		if head == temp {
			break
		}
	}
	fmt.Printf("last boy:%d out\n", head.no)
}
