package entities

// Ether is the main usage of a 'native' currency, i.e. for Ethereum mainnet and all testnets
type Ether struct {
	*BaseCurrency
}

func EtherOnChain(chainId uint) *Ether {
	ether := &Ether{
		BaseCurrency: &BaseCurrency{
			IsNativez: true,
			IsTokenz:  false,
			ChainIdz:  chainId,
			Decimalsz: 18,
			Symbolz:   "ETH",
			Namez:     "Ether",
		},
	}
	ether.BaseCurrency.Currency = ether
	return ether
}

func (e *Ether) Equal(other Currency) bool {
	v, isEther := other.(*Ether)
	if isEther {
		return v.IsNativez && v.ChainIdz == e.ChainIdz

	}
	return false
}

func (e *Ether) Wrapped() *Token {
	return WETH9[e.ChainId()]
}
