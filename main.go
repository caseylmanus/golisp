package main

import (
	"fmt"
	"log"
)

func main() {
	exp, err := ParseExpression("(+ 3 5 (+ 3 3 9999) (+ 8 -998) (+ 88 999))")
	if err != nil {
		log.Fatal(err)
	}
	v, err := exp.Eval()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

}
