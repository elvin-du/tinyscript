package lexer

import "regexp"

var (
	ptnLetter   = regexp.MustCompile("^[a-zA-Z]$")
	ptnNumber   = regexp.MustCompile("^[0-9]$")
	ptnLiteral  = regexp.MustCompile("^[_a-zA-Z0-9]$")
	ptnOperator = regexp.MustCompile("^[+-\\\\*<>=!&|^%/]$")
)

func IsLetter(c string) bool {
	return ptnLetter.MatchString(c)
}

func IsNumber(c string) bool {
	return ptnNumber.MatchString(c)
}

func IsLiteral(c string) bool {
	return ptnLiteral.MatchString(c)
}

func IsOperator(c string) bool {
	return ptnOperator.MatchString(c)
}
