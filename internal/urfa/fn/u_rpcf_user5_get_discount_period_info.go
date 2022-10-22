package fn

// rpcf_user5_get_discount_period_info
func xU402e(c conn, p Dict) Dict {
	c.Call(-0x402e)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":                      c.GetI(),
		"begin":                   c.GetI(),
		"end":                     c.GetI(),
		"periodic_type":           c.GetI(),
		"next_discount_period_id": c.GetI(),
		"discount_interval":       c.GetI(),
		"canonical_len":           c.GetI(),
		"custom_duration":         c.GetI(),
		"static_id":               c.GetI(),
	}
}
