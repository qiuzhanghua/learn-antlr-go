package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"learn-antlr-go/calc"
	"os"
)

//go:generate antlr4 -Dlanguage=Go -visitor -package calc -o calc Expr.g4

type TreeShapeListener struct {
	*calc.BaseExprListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := calc.NewExprLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := calc.NewExprParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
