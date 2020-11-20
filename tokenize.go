package main

import "unicode"

func tokenize(text string) []string {
	token := ""
	var tokens []string
	openParens := 0
	ticked := false
	quoted := false
	lastPos := 0
	for i, r := range text {
		s := string(r)
		switch {
		case !quoted && s == "`" && string(text[lastPos]) != `\`:
			ticked = !ticked
			token = token + s
		case !ticked && s == `"` && string(text[lastPos]) != `\`:
			quoted = !quoted
			token = token + s
		case !quoted && !ticked && s == "(":
			token = token + s
			openParens++
		case !quoted && !ticked && s == ")":
			token = token + s
			openParens--
			if openParens == 0 {
				tokens = append(tokens, token)
				token = ""
			}
		case !ticked && !quoted && unicode.IsSpace(r) && openParens == 0:
			if len(token) > 0 {
				tokens = append(tokens, token)
				token = ""
			}
		default:
			token = token + s
		}
		lastPos = i
	}
	if len(token) > 0 {
		tokens = append(tokens, token)
	}
	return tokens
}
