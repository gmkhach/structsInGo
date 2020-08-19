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
		// field:value initialization
		p1 := person{
			first: "James",
			last: "Bond",
			age: 32,
		}

		// value initialization
		p2 := person{
			"Miss",
			"Moneypenny",
			27,
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
		
		// you can access values of embeded types both by treating them as promoted values  or by the full path
		fmt.Println(sa1.first, sa1.last, sa1.person.age, sa1.ltk)

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