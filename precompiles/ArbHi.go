package precompiles

// ArbHi provides a friendly greeting to anyone who calls it.
type ArbHi struct {
	Address addr // 0x11a, for example
}

func (con *ArbHi) SayHi(c ctx, evm mech) (string, error) {
	return "hi", nil
}

func (con *ArbHi) GetNumber(c ctx, evm mech) (uint64, error) {
	return c.State.GetMyNumber()
}

func (con *ArbHi) SetNumber(c ctx, evm mech, newNumber uint64) error {
	return c.State.SetNewMyNumber(newNumber)
}
