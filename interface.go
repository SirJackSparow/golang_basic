package main

type HashName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHellos(hashName HashName) {
	println("Hello", hashName.GetName())
}

func main() {

	var namePerson Person
	namePerson.Name = "Maulana"
	SayHellos(namePerson)

}
