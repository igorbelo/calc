package main

import (
	"fmt"
	"github.com/igorbelo/gocalc/syntax"
)

func main() {
	lexer := syntax.NewLexer("123 + 3 / 2\n1+2")

	for {
		token := lexer.Lex()
		if token == nil {
			break
		}
		fmt.Println(token.ID)
	}
}
