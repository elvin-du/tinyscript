package lexer

var KeyWords = map[string]bool{
	"var":    true,
	"if":     true,
	"else":   true,
	"for":    true,
	"while":  true,
	"break":  true,
	"func":   true,
	"return": true,
}

func IsKeyword(key string) bool {
	return KeyWords[key]
}
