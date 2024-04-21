package main

import "fmt"

type Person struct {
	Name       string
	Age        int
	Occupation string
}

func (p *Person) ChangeOccupation(newOccupation string) {
	p.Occupation = newOccupation
}

func (p *Person) ChangeName(newName string) {
	p.Name = newName
}

func (p *Person) ChangeAge(newAge int) {
	p.Age = newAge
}

func main() {
	age := 30
	fmt.Println("Age before:", age)
	alterAge(&age)
	fmt.Println("Age after:", age)

	person := Person{
		Name:       "Elijah Hallan",
		Age:        30,
		Occupation: "Software Engineer",
	}

	fmt.Println("Name before:", person.Name)
	person.ChangeName("Ezra Hallan")
	fmt.Println("Name after:", person.Name)

	fmt.Println("Age before:", person.Age)
	person.ChangeAge(25)
	fmt.Println("Age after:", person.Age)
}

// alterAge is a function that takes a pointer to an integer as an argument.
// It changes the value of the integer pointed to by the pointer to 25.
// It then returns the new value.
func alterAge(age *int) int {
	*age = 25
	return *age
}
