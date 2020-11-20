package main

import (
	"fmt"
	"log"
)

func main() {
	exp, err := ParseExpression("(* (* 4 5) 8 (* 4 5))")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exp)

}
