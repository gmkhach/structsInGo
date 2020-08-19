package main

import "fmt"

// Structs allow to aggregate elements of same or different types together.
// The names of the elements are called fields and can be defined explicitly (identified elements) or implicitely (anonymous fields).
type person struct {
	first 	string
	last	string
	age		int
}

type secretAgent struct {
	person	
	ltk 	bool
}
func main() {
		p1 := person{
			first: "James",
			last: "Bond",
			age: 32,
		}

		p2 := person{
			first: "Miss",
			last: "Moneypenny",
			age: 27,
		}

		fmt.Println(p1)
		fmt.Println(p2)
		
		fmt.Println(p1.first, p1.last, p1.age)
		fmt.Println(p1.first, p2.last, p2.age)

		// You can embed structs into other structs
		// This is called an anonymous field, or an embeded type
		sa1 := secretAgent{
			person: person{
				first: "James",
				last: "Bond",
				age: 32,
			},
			ltk: true,
		}
		
		fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

		// Creating an anonymous struct
		p := struct {
			first	string
			last 	string
			age 	int
		}{
			first: "Dr",
			last: "No",
			age: 40,
		}

		fmt.Println(p)
}