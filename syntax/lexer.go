package syntax

import (
	"regexp"
)

type rule struct {
	id    string
	regex string
}

var rules []*rule = []*rule{
	&rule{"integer", "^[0-9]+$"},
	&rule{"operator", "^[+\\-\\*\\/]$"},
	&rule{"endl", "^\n$"},
	&rule{"whtspc", "^ $"},
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

func (l *lexer) logMatch() *Token {
	if l.match.length == 0 {
		return &Token{ID: "unknown"}
	}

	token := &Token{ID: l.match.rule.id}
	l.clearMatch()

	return token
}

func (l *lexer) clearMatch() {
	l.buffer = ""
	l.cursor = l.match.cursor
	l.match = match{length: 0}
}

func Lex(input string) func() *Token {
	l := lexer{cursor: -1, input: input}

	return func() *Token {
		var token *Token

		for l.moveCursor() {
			l.addToBuffer()
			l.findMatch()
			if !l.hasNextChar() {
				token = l.logMatch()
				break
			}
		}

		return token
	}
}
