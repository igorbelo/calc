package syntax

import (
	"fmt"
	"regexp"
)

type rule struct {
	id    string
	regex string
}

var rules []*rule = []*rule{
  &rule{"integer",  "^[0-9]+$"},
  &rule{"operator", "^[+\\-\\*\\/]$"},
  &rule{"endl",     "^[;\n]$"},
  &rule{"whtspc",   "^ $"},
  &rule{"weird",    "^_a$"},
}

type match struct {
	cursor int
	length int
	rule   *rule
}

type Token struct {
	ID string
}

type lexer struct {
	cursor int
	match  match
	input  string
	buffer string
	tokens []Token
}

func (l *lexer) moveCursor() bool {
	if !l.hasNextChar() {
		return false
	}

	l.cursor += 1
	return true
}

func (l *lexer) hasNextChar() bool {
	return l.cursor+1 < len(l.input)
}

func (l *lexer) currentChar() byte {
	return l.input[l.cursor]
}

func (l *lexer) addToBuffer() {
	l.buffer += string(l.currentChar())
}

func (l *lexer) findMatch() {
	for _, rule := range rules {
		if hasMatch, _ := regexp.MatchString(rule.regex, l.buffer); hasMatch {
			l.match = match{cursor: l.cursor, length: len(l.buffer), rule: rule}
			break
		}
	}
}

func (l *lexer) logMatch() {
  if l.match.length == 0 {
    fmt.Println("Unknown token")
  } else {
  	l.tokens = append(l.tokens, Token{ID: l.match.rule.id})
    l.clearMatch();
  }
}

func (l *lexer) clearMatch() {
	l.buffer = ""
	l.cursor = l.match.cursor
	l.match  = match{length: 0}
}

func Lex(input string) []Token {
	l := lexer{cursor: -1, input: input}

	for l.moveCursor() {
		l.addToBuffer()
		l.findMatch()
		if !l.hasNextChar() {
			l.logMatch()
		}
	}

	return l.tokens
}
