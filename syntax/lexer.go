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

type Lexer struct {
	cursor int
	match  match
	input  string
	buffer string
}

func (l *Lexer) moveCursor() bool {
	if !l.hasNextChar() {
		return false
	}

	l.cursor += 1
	return true
}

func (l *Lexer) hasNextChar() bool {
	return l.cursor+1 < len(l.input)
}

func (l *Lexer) currentChar() byte {
	return l.input[l.cursor]
}

func (l *Lexer) addToBuffer() {
	l.buffer += string(l.currentChar())
}

func (l *Lexer) findMatch() {
	for _, rule := range rules {
		if hasMatch, _ := regexp.MatchString(rule.regex, l.buffer); hasMatch {
			l.match = match{cursor: l.cursor, length: len(l.buffer), rule: rule}
			break
		}
	}
}

func (l *Lexer) logMatch() *Token {
	if l.match.length == 0 {
		return &Token{ID: "unknown"}
	}

	token := &Token{ID: l.match.rule.id}
	l.clearMatch()

	return token
}

func (l *Lexer) clearMatch() {
	l.buffer = ""
	l.cursor = l.match.cursor
	l.match = match{length: 0}
}

func NewLexer(input string) *Lexer {
	return &Lexer{cursor: -1, input: input}
}

func (l *Lexer) Lex() *Token {
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
