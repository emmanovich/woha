package main

type Dataset struct{}

func (d Dataset) Build() []Person {

	gen := Generator{}

	list := []Person{}

	for i := 0; i < RecordCount; i++ {

		list = append(

			list,

			gen.Create(),

		)

	}

	return list

}
