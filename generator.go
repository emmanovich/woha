package main

type Generator struct{}

func (g Generator) Create() Person {

	return Person{

		Name: randomName(),

		Address: randomAddress(),

		Phone: randomPhone(),

		Date: randomDate(),

	}

}
