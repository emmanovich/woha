package main

import "fmt"

func ExportCSV(

	data []Person,

) {

	fmt.Println()

	fmt.Println("CSV Output")

	fmt.Println()

	fmt.Println(

		"name,address,phone,date",

	)

	for _, p := range data {

		fmt.Printf(

			"%s,%s,%s,%s\n",

			p.Name,

			p.Address,

			p.Phone,

			p.Date,

		)

	}

}
