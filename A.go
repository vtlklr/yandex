package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type customStack struct {
	stack []int
	lock  sync.RWMutex
}

func (c *customStack) Push(name int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customStack) Pop() error {
	l := len(c.stack)
	if l > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:l-1]
		return nil
	}
	return fmt.Errorf("pop Error: Stack is empty")
}

func (c *customStack) Front() (int, error) {
	l := len(c.stack)
	if l > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[l-1], nil
	}
	return 0, fmt.Errorf("peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return len(c.stack)
}

func (c *customStack) Empty() bool {
	return len(c.stack) == 0
}
func (c *customStack) Max() int {
	var max int
	for i := 0; i < len(c.stack); i++ {
		if c.stack[i] >= max {
			max = c.stack[i]
		}
	}
	return max
}
func Scan() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	return str
}
func main() {

	m := Scan()
	m = strings.TrimSpace(m)
	n, err := strconv.Atoi(m)
	if err != nil {
		fmt.Println(err)
	}

	comm := make([]string, n)

	for i := 0; i < n; i++ {
		comm[i] = strings.TrimSpace(Scan())
	}
	customStack := &customStack{
		stack: make([]int, 0),
	}

	for i := 0; i < n; i++ {
		switch comm[i] {
		case "max":
			{
				fmt.Println(customStack.Max())
			}
		case "pop":
			{
				err := customStack.Pop()
				if err != nil {
					return
				}
			}
		default:
			{
				s1 := strings.Split(comm[i], " ")
				d, _ := strconv.Atoi(s1[1])
				customStack.Push(d)
			}

		}
	}
	/*
		fmt.Printf("Push: A\n")
		customStack.Push(2)
		fmt.Printf("Push: B\n")
		customStack.Push(3)
		fmt.Printf("Size: %d\n", customStack.Size())
		for customStack.Size() > 0 {
			fmt.Println("max: ", customStack.Max())
			frontVal, _ := customStack.Front()
			fmt.Printf("Front: %d\n", frontVal)
			fmt.Printf("Pop: %d\n", frontVal)
			customStack.Pop()
		}
		fmt.Printf("Size: %d\n", customStack.Size())
	*/
}
