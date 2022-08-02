package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type cabs struct {
	driver int
	rw     sync.RWMutex
	wg     sync.WaitGroup
}

func (c *cabs) bookCab(name string) {
	defer c.wg.Done()
	c.rw.Lock()
	fmt.Println("welcome to cab service", name)
	if c.driver >= 1 {
		fmt.Println("car is available", name)
		time.Sleep(3 * time.Second)
		fmt.Println("booking confirmed", name)
		c.driver--
	} else {
		fmt.Println("car is not avilable", name)
	}
	c.rw.Unlock()
}
func (c *cabs) getCabDriver() {
	defer c.wg.Done()
	c.rw.Lock()
	fmt.Println("driver", c.driver)
	c.rw.Unlock()
}
func main() {
	c := cabs{
		driver: 5,
	}
	for i := 1; i <= 5; i++ {
		c.wg.Add(1)
		go c.getCabDriver()
	}
	for i := 1; i <= 15; i++ {
		c.wg.Add(1)
		go c.bookCab("user " + strconv.Itoa(i))
	}
	for i := 1; i <= 15; i++ {
		c.wg.Add(1)
		go c.getCabDriver()
	}
	c.wg.Wait()
}
