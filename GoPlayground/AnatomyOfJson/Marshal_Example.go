package main
import (
	"encoding/json"
	"log"
)
type Employee struct{
	Name string
	Age int 
	Salary int 
}
func MarshallingExample(){
	empObj := Employee{
		Name:"Vibhor",
		Age :24,
		Salary: 344334,
	}
	empResponse , err := json.Marshal(empObj)
	if err != nil {
		log.Println("Error while marshalling JSON:",err)
	}
	log.Println("Marshalling Response : ",string(empResponse))
}

func main(){
	MarshallingExample()
}