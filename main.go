package main

import (
	"fmt"
	"github.com/igorbelo/gocalc/syntax"
)

func main() {
	lex := syntax.Lex("123 + 3 / 2\n1+2")

	for {
		token := lex()
		if token == nil {
			break
		}
		fmt.Println(token.ID)
	}
}
