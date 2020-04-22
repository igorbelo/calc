package main

import (
	"github.com/igorbelo/gocalc/syntax"
	"fmt"
)

func main() {
	tokens := syntax.Lex("123 _a +  3 / 2 ")

	for _, token := range tokens {
		fmt.Println(token.ID)
	}
}
