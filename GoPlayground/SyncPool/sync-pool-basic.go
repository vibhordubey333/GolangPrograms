
// A dummy struct
type Person struct {
	Name string
}

// Initializing pool
var personPool = sync.Pool{
	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	New: func() interface{} { return new(Person) },
}

// Main function
func main() {
	// Get hold of an instance
	newPerson := personPool.Get().(*Person)
	// Defer release function
	// After that the same instance is 
	// reusable by another routine
	defer personPool.Put(newPerson)

	// Using the instance
	newPerson.Name = "Jack"
}
