package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		exp, err := ParseExpression(scanner.Text())
		if err != nil {
			log.Println(err)
			continue
		}
		v, err := exp.Eval()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(v)
	}

}
