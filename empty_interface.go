package main

func Ups(s string) interface{} {
	return s
}

func main() {
	var data interface{} = Ups("Benar")
	print(data)
}
