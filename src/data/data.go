package data

import "fmt"

type User struct {
	Name string
	Age  uint
}

func getNewInstance() *User {
	return new(User)
}

func getInstance() *User {
	return &User{}
}

func InstanceTest() {
	a := getNewInstance()
	fmt.Println("getNewInstance().")
	fmt.Printf("User.name: %s\n", a.Name)

	b := getInstance()
	fmt.Println("getInstance().")
	fmt.Printf("User.name: %s\n", b.Name)
}
