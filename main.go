package main

import (
	"fmt"
	"time"
)

type User struct {
	A int32
	B []int32
	C string
	D bool
	E struct{}
}
type myMutex chan struct{}

func NewMutex() myMutex {
	ch := make(myMutex, 1)
	return ch
}
func (m *myMutex) Lock() bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case *m <- struct{}{}:
		return true
	case <-timer.C:
		fmt.Println("Time out...")
		return false
	}
}
func (m *myMutex) UnLock() {
	select {
	case <-*m:
	}
}
func add(a *int, m *myMutex) {
	m.Lock()
	*a++
	m.UnLock()
}
func main() {
	//mutex := NewMutex()
	//j := 0
	//for i := 0; i < 50; i++ {
	//	go add(&j, &mutex)
	//}
	//time.Sleep(time.Second)
	//fmt.Println(j)
	DoServer("8888")
}
