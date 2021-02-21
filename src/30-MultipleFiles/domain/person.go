package domain

type Person struct {
	Name string
	Age  int
}

func (person *Person) SetName(name string) {
	person.Name = name
}
