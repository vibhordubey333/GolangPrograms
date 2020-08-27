package main

import (
	"fmt"
	//"runtime"
	"sync"
	"time"
)

type input struct {
	no  int
	mux sync.Mutex
}

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	in := input{no: 2}
	go in.Cube(&in)
	go in.Square(&in)
	time.Sleep(time.Second * 1)

	wg.Wait()
}

func (obj *input) Square(in *input) {
	obj.mux.Lock()
	defer obj.mux.Unlock()
	obj.no = obj.no * obj.no
	fmt.Println("Square: ", obj.no)
	wg.Done()
}

func (obj *input) Cube(in *input) {
	obj.mux.Lock()
	defer obj.mux.Unlock()
	obj.no = obj.no * obj.no * obj.no
	fmt.Println("Cube: ", obj.no)
	wg.Done()
}
