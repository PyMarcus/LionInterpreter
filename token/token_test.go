package token

import "testing"

func TestToken(t *testing.T) {
	tok := &Token{Type: "string", Literal: "o"}
	if tok.Type != "string" && tok.Literal != "o" {
		t.Fatalf("tests - tokentype wrong. expected=%q, got=%q",
			"string", tok.Type)
	}
}
