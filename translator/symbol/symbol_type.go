package symbol

// SymbolType 符号类型
type SymbolType string

const (
	ADDRESS_SYMBOL   SymbolType = "address symbol"
	IMMEDIATE_SYMBOL SymbolType = "immediate symbol"
	LABEL_SYMBOL     SymbolType = "label symbol"
)
