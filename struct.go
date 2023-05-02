package main

// prototype data
type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	jumlah := Customer{
		Name:    "KG",
		Address: "People",
		Age:     12,
	}

	print(jumlah)

	gunung := Customer{"Kembar", "Ganjil", 10}

	print(gunung)

	var bima Customer
	bima.Name = "Bima"
	bima.Address = "Indonesia"
	bima.Age = 12
	bima.sayHello()

}

//struct funtion

func (customer Customer) sayHello() {
	println("Hello: ", customer.Address)
}
