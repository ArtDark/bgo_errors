package main

import (
	"fmt"
	"github.com/ArtDark/bgo_errors/pkg/card"
	"github.com/ArtDark/bgo_errors/pkg/transfer"
)

func main() {

	svc := card.New("YourBank")

	visa := svc.CardIssue(
		"0001",
		"Artem",
		"Balusov",
		"Visa",
		12345_67,
		"RUR",
		"5106 2142 2342 4322",
	)

	master := svc.CardIssue(
		"0002",
		"Ivan",
		"Ivanov",
		"MasterCard",
		98765_43,
		"RUR",
		"5106 2142 4322 2342",
	)

	fmt.Println(svc)
	fmt.Println(visa)
	fmt.Println(master)

	s := transfer.IsValid("4  5  6  1     2  6  1  2     1  2  3  4     5  4  6  7")

	fmt.Println(s)

}
