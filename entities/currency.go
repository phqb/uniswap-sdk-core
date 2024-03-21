package entities

// Currency is any fungible financial instrument, including Ether, all ERC20 tokens, and other chain-native currencies
type Currency interface {
	IsNative() bool
	IsToken() bool
	ChainId() uint
	Decimals() uint
	Symbol() string
	Name() string
	Equal(other Currency) bool
	Wrapped() *Token
}

// BaseCurrency is an abstract struct, do not use it directly
type BaseCurrency struct {
	Currency  Currency
	IsNativez bool   // Returns whether the currency is native to the chain and must be wrapped (e.g. Ether)
	IsTokenz  bool   // Returns whether the currency is a token that is usable in Uniswap without wrapping
	ChainIdz  uint   // The chain ID on which this currency resides
	Decimalsz uint   // The decimals used in representing currency amounts
	Symbolz   string // The symbol of the currency, i.e. a short textual non-unique identifier
	Namez     string // The name of the currency, i.e. a descriptive textual non-unique identifier
}

func (c *BaseCurrency) IsNative() bool {
	return c.IsNativez
}

func (c *BaseCurrency) IsToken() bool {
	return c.IsTokenz
}

func (c *BaseCurrency) ChainId() uint {
	return c.ChainIdz
}

func (c *BaseCurrency) Decimals() uint {
	return c.Decimalsz
}

func (c *BaseCurrency) Symbol() string {
	return c.Symbolz
}

func (c *BaseCurrency) Name() string {
	return c.Namez
}

// Equal returns whether the currency is equal to the other currency
func (c *BaseCurrency) Equal(other Currency) bool {
	panic("Equal method has to be overridden")
}

func (c *BaseCurrency) Wrapped() *Token {
	panic("Wrapped method has to be overridden")
}
