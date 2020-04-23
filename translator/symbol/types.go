package symbol

type SymbolType int

const (
	SYMBOL_ADDRESS SymbolType = iota
	SYMBOL_IMMEDIATE
	SYMBOL_LABEL
)

func (s SymbolType) String() string {
	switch s {
	case SYMBOL_ADDRESS:
		return "symbol_address"
	case SYMBOL_IMMEDIATE:
		return "symbol_immediate"
	case SYMBOL_LABEL:
		return "symbol_label"
	}

	panic("unknown symbol type")
}
