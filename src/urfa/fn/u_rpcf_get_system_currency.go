package fn

// rpcf_get_system_currency
func x0101(c conn, _ Dict) Dict {
	c.Call(0x0101)

	return Dict{
		"currency_id": c.GetI(),
	}
}
