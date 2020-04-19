package symbol

type SymbolType int

const (
	SYMBOL_ADDRESS SymbolType = iota
	SYMBOL_IMMEDIATE
	SYMBOL_LABEL
)
