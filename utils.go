package main

import (
	"fmt"
	"math/rand"
)

func randomName() string {

	return Names[rand.Intn(len(Names))]

}

func randomAddress() string {

	number := rand.Intn(900) + 100

	return fmt.Sprintf(

		"%d %s",

		number,

		Streets[rand.Intn(len(Streets))],

	)

}

func randomPhone() string {

	return fmt.Sprintf(

		"+1-555-%04d",

		rand.Intn(10000),

	)

}

func randomDate() string {

	day := rand.Intn(28) + 1

	month := rand.Intn(12) + 1

	return fmt.Sprintf(

		"2024-%02d-%02d",

		month,

		day,

	)

}
