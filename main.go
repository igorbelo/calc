package main

import (
  "fmt"
  "strconv"

  "github.com/igorbelo/gocalc/parser"
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/llir/llvm/ir"
  "github.com/llir/llvm/ir/constant"
  "github.com/llir/llvm/ir/types"
  "github.com/llir/llvm/ir/value"
)

var i64 = types.I64
var m = ir.NewModule()
var fun = m.NewFunc("calc", i64)
var block = fun.NewBlock("")

type calcListener struct {
  *parser.BaseCalcListener

  stack []value.Value
}

func (l *calcListener) push(v value.Value) {
  l.stack = append(l.stack, v)
}

func (l *calcListener) pop() value.Value {
  if len(l.stack) < 1 {
    panic("stack is empty unable to pop")
  }

  // Get the last value from the stack.
  result := l.stack[len(l.stack)-1]

  // Pop the last element from the stack.
  l.stack = l.stack[:len(l.stack)-1]

  return result
}

// ExitMulDiv is called when exiting the MulDiv production.
func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
  right, left := l.pop(), l.pop()

  switch c.GetOp().GetTokenType() {
  case parser.CalcParserMUL:
    l.push(block.NewMul(left, right))
  case parser.CalcParserDIV:
    l.push(block.NewUDiv(left, right))
  default:
    panic(fmt.Sprintf("unexpected operation: %s", c.GetOp().GetText()))
  }
}

// ExitAddSub is called when exiting the AddSub production.
func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
  right, left := l.pop(), l.pop()

  switch c.GetOp().GetTokenType() {
  case parser.CalcParserADD:
    l.push(block.NewAdd(left, right))
  case parser.CalcParserSUB:
    l.push(block.NewSub(left, right))
  default:
    panic(fmt.Sprintf("unexpected operation: %s", c.GetOp().GetText()))
  }
}

// ExitNumber is called when exiting the Number production.
func (l *calcListener) ExitNumber(c *parser.NumberContext) {
  i, err := strconv.Atoi(c.GetText())
  if err != nil {
    panic(err.Error())
  }

  l.push(constant.NewInt(i64, int64(i)))
}

func (l *calcListener) ExitStart(_ *parser.StartContext) {
  block.NewRet(l.pop())
}

// calc takes a string expression and returns the evaluated result.
func calc(input string) {
  // Setup the input
  is := antlr.NewInputStream(input)

  // Create the Lexer
  lexer := parser.NewCalcLexer(is)
  stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

  // Create the Parser
  p := parser.NewCalcParser(stream)

  // Finally parse the expression (by walking the tree)
  var listener calcListener
  antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())
}

func main() {
  calc("1 + 2 * 3")
  fmt.Println(m)
}
