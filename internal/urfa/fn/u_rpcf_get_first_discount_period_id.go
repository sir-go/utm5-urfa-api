package fn

// rpcf_get_first_discount_period_id
func x2601(c conn, _ Dict) Dict {
	c.Call(0x2601)

	return Dict{
		"id": c.GetI(),
	}
}
