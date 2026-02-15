package parser

import (
	"github.com/PyMarcus/lioninterpreter/ast"
	"github.com/PyMarcus/lioninterpreter/lexer"
	"github.com/PyMarcus/lioninterpreter/token"
)

/*Dá significado aos tokens identificados pelo lexer, ou seja, coloca em ordem gramatical.*/

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *ast.Program {
	return nil
}
