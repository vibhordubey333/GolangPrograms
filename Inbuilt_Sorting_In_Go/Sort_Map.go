package main
import (
	"fmt"
	"sort"
)
func main(){
	m := map[string]int{
		"Z" : 1,
		"A" : 2,
		"C" : 3,
	} 
	keys := make([]string,0,len(m))
	for k := range m {
		keys = append(keys,k)
	}
	sort.Strings(keys)

	for _ , k := range keys{
		fmt.Println(k,m[k])
	}

}
