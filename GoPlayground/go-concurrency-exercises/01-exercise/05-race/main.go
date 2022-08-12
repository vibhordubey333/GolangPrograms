//------------- Program With Race Condition ------------------------------
/*package main
import(
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
	}
	wg.Wait()
}
*/

//------------ Program with Race Condition Fixed -------------------------

package main
import(
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(j int) {
			fmt.Println(j) // Read local copy of loop counter.
			wg.Done()
		}(i)
	}
	wg.Wait()
}
