package syntax

import (
	"regexp"
)

type Token int

const (
	UNKNOWN Token = iota
	INTEGER
	OPERATOR
	ENDL
	WHTSPC
)

type rule struct {
	token Token
	regex string
}

var rules []*rule = []*rule{
	&rule{INTEGER,  "^[0-9]+$"},
	&rule{OPERATOR, "^[+\\-\\*\\/]$"},
	&rule{ENDL,     "^\n$"},
	&rule{WHTSPC,   "^ $"},
}

type match struct {
	cursor int
	buffer string
	rule   *rule
}

func (m *match) length() int {
	return len(m.buffer)
}

type Lexer struct {
	cursor int
	match  match
	input  string
	buffer string
}

func (l *Lexer) moveCursor() bool {
	if !l.CanLex() {
		return false
	}

	l.cursor += 1
	return true
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
			l.match = match{cursor: l.cursor, buffer: l.buffer, rule: rule}
			break
		}
	}
}

func (l *Lexer) logMatch() (Token, string) {
	if l.match.length() == 0 {
		return UNKNOWN, l.match.buffer
	}

	return l.match.rule.token, l.match.buffer
}

func (l *Lexer) clearMatch() {
	l.buffer = ""
	l.cursor = l.match.cursor
	l.match = match{buffer: ""}
}

func NewLexer(input string) *Lexer {
	return &Lexer{cursor: -1, input: input}
}

func (l *Lexer) CanLex() bool {
	return l.cursor+1 < len(l.input)
}

func (l *Lexer) Lex() (Token, string) {
	var token Token
	var buffer string

	for l.moveCursor() {
		l.addToBuffer()
		l.findMatch()
		if !l.CanLex() {
			token, buffer = l.logMatch()
			l.clearMatch()
			break
		}
	}

	return token, buffer
}
