package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(

		time.Now().UnixNano(),

	)

	fmt.Println(

		"Generating records...\n",

	)

	data := Dataset{}.

		Build()

	fmt.Printf(

		"Records created: %d\n\n",

		len(data),

	)

	for _, person := range data {

		fmt.Println(

			person.Name,

		)

		fmt.Println(

			person.Address,

		)

		fmt.Println(

			person.Phone,

		)

		fmt.Println(

			person.Date,

		)

		fmt.Println()

	}

	Export(

		data,

	)

	fmt.Println()

	fmt.Println(

		"Export completed.",

	)

}
