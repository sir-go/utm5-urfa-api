package fn

// rpcf_dealer_get_system_currency
func x7000003f(c conn, _ Dict) Dict {
	c.Call(0x7000003f)

	return Dict{
		"currency_id": c.GetI(),
	}
}
