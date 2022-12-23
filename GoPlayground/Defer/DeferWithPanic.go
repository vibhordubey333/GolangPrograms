// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		time.Sleep(5 * time.Second)
		fmt.Println("testing")
	}()
	panic(0)
}
/*
Output:
testing
panic: 0

goroutine 1 [running]:
main.main()
	/tmp/sandbox2365228083/prog.go:15 +0x49

Program exited.
*/
