package main

import (
	"encoding/json"
	"fmt"
)

func ExportJSON(

	data []Person,

) {

	bytes, _ := json.MarshalIndent(

		data,

		"",

		"  ",

	)

	fmt.Println()

	fmt.Println("JSON Output")

	fmt.Println()

	fmt.Println(string(bytes))

}
