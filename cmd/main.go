package main

import (
	"fmt"
	"os"
	"strings"

	ub64c "github.com/Blackmocca/uuid-base64-converter"
)

func main() {
	input := strings.Trim(os.Args[1], " ")

	uid, b64, err := ub64c.Convert(input)
	if err != nil {
		panic(err)
	}

	mongoData := ub64c.NewMongoBinary(b64)

	fmt.Println("uid		: ", uid)
	fmt.Println("base64		: ", b64)
	fmt.Println("mongoBinaryData	: ", mongoData.ToJSON())
}
