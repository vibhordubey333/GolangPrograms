package main

import "fmt"

/*
The Facade pattern in Go provides a simplified interface to a complex subsystem, making it easier to use. It hides the complexities of the subsystem and provides a higher-level interface for the client to interact with. Here's an example:
Let's say we have a complex subsystem with multiple components:
*/

// SubSystemA represents a complex subsystem.
type SubSystemA struct{}

func (s *SubSystemA) operationA() string {
	return "Subsystem A: Ready!"
}

// SubSystemB represents another complex subsystem.
type SubSystemB struct{}

func (s *SubSystemB) operationB() string {
	return "Subsystem B: Go!"
}

// Facade provides a simplified interface to the complex subsystem.
type Facade struct {
	subSystemA *SubSystemA
	subSystemB *SubSystemB
}

func NewFacade() *Facade {
	return &Facade{
		subSystemA: &SubSystemA{},
		subSystemB: &SubSystemB{},
	}
}

func (f *Facade) PerformOperation() string {
	result := "Facade initializes subsystems:\n"
	result += f.subSystemA.operationA() + "\n"
	result += f.subSystemB.operationB()
	return result
}

func main() {
	// Using the facade to perform the operation without dealing with subsystem complexities directly
	facade := NewFacade()
	result := facade.PerformOperation()
	fmt.Println(result)
}

/*
Facade initializes subsystems:
Subsystem A: Ready!
Subsystem B: Go!

*/
