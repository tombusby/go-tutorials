package main

func log(message string) {

}

// both variables have same type:
//  can use `a, b int` instead of `a int, b int`
func add(a, b int) int {
	return a + b
}

func power(name string) (int, bool) {
	return 0, false
}

// named return values, note = instead of := (already declared)
func power2(name string) (power int, exists bool) {
	power, exists = 9000, true
	return
}

func main() {
	// Explicit sub-scope
	{
		// Getting both values from power()
		value, exists := power("goku")
		if exists == false {
			println(add(value, 20))
		}
	}

	{
		// Getting one values from power() and discarding the first
		_, exists := power2("goku")
		if exists == false {
			// handle error
		}

		// using = instead of := because all vars already declared
		_, exists = power("someone else")
		if exists == false {
			// handle error
		}
	}
}
