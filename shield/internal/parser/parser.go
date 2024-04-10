package parser

import (
	"fmt"
	"strconv"

	"github.com/warden-protocol/wardenprotocol/shield/ast"
	"github.com/warden-protocol/wardenprotocol/shield/internal/lexer"
	"github.com/warden-protocol/wardenprotocol/shield/token"
)

const (
	_ int = iota
	LOWEST
	OR
	AND
	CALL
)

var precedences = map[token.Type]int{
	token.Type_OR:     OR,
	token.Type_AND:    AND,
	token.Type_LPAREN: CALL,
}

type (
	prefixParseFn func() *ast.Expression
	infixParseFn  func(*ast.Expression) *ast.Expression
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.Type]prefixParseFn
	infixParseFns  map[token.Type]infixParseFn
	errors         []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:              l,
		prefixParseFns: make(map[token.Type]prefixParseFn),
		infixParseFns:  make(map[token.Type]infixParseFn),
	}

	p.registerPrefix(token.Type_IDENT, p.parseIdentifier)
	p.registerPrefix(token.Type_INT, p.parseIntegerLiteral)
	p.registerPrefix(token.Type_TRUE, p.parseBooleanLiteral)
	p.registerPrefix(token.Type_FALSE, p.parseBooleanLiteral)
	p.registerPrefix(token.Type_LBRACKET, p.parseArrayLiteral)

	p.registerInfix(token.Type_AND, p.parseInfixExpression)
	p.registerInfix(token.Type_OR, p.parseInfixExpression)
	p.registerInfix(token.Type_LPAREN, p.parseCallExpression)

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) Parse() *ast.Expression {
	return p.parseExpression(LOWEST)
}

func (p *Parser) parseExpression(precedence int) *ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.Type_SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseIdentifier() *ast.Expression {
	return ast.NewIdentifier(&ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	})
}

func (p *Parser) parseIntegerLiteral() *ast.Expression {
	v, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err != nil {
		msg := "could not parse %q as integer"
		p.errors = append(p.errors, fmt.Sprintf(msg, p.curToken.Literal))
		return nil
	}
	return ast.NewIntegerLiteral(&ast.IntegerLiteral{
		Token: p.curToken,
		Value: v,
	})
}

func (p *Parser) parseBooleanLiteral() *ast.Expression {
	switch p.curToken.Type {
	case token.Type_TRUE:
		return ast.NewBooleanLiteral(&ast.BooleanLiteral{Token: p.curToken, Value: true})
	case token.Type_FALSE:
		return ast.NewBooleanLiteral(&ast.BooleanLiteral{Token: p.curToken, Value: false})
	default:
		p.errors = append(p.errors, "expected true or false")
		return nil
	}
}

func (p *Parser) parseArrayLiteral() *ast.Expression {
	al := &ast.ArrayLiteral{Token: p.curToken}
	al.Elements = p.parseExpressionList(token.Type_RBRACKET)
	return ast.NewArrayLiteral(al)
}

func (p *Parser) parseInfixExpression(left *ast.Expression) *ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)
	return ast.NewInfixExpression(exp)
}

func (p *Parser) parseCallExpression(function *ast.Expression) *ast.Expression {
	ident, ok := ast.UnwrapIdentifier(function)
	if !ok {
		p.errors = append(p.errors, "expected identifier")
		return nil
	}

	exp := &ast.CallExpression{Token: p.curToken, Function: ident}
	exp.Arguments = p.parseExpressionList(token.Type_RPAREN)
	return ast.NewCallExpression(exp)
}

func (p *Parser) parseExpressionList(end token.Type) []*ast.Expression {
	args := []*ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.Type_COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return args
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekError(t token.Type) {
	msg := "expected next token to be %s, got %s instead"
	p.errors = append(p.errors, fmt.Sprintf(msg, t, p.peekToken.Type))
}

func (p *Parser) Errors() []string {
	return p.errors
}
