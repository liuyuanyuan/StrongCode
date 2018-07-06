package lib

type Person struct {
	Age  int
	Name string
}

func (p Person) GetName() string {
	return p.Name
}
